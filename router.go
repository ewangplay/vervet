package vervet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

type Router struct {
	routes map[string]Handler
}

func NewRouter(routes map[string]Handler) *Router {
	return &Router{routes}
}

func (this *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//路由分发设置，用来判断url是否合法，通过配置文件的正则表达式配置

	header := w.Header()
	header.Add("Content-Type", "application/json")
	header.Add("charset", "UTF-8")

	version, resources, err := this.ParseURL(r.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, MakeErrorResult(-1, err.Error()))
		return
	}

	var res []string
	var resource string
	for i, r := range resources {
		if i == 0 {
			resource = r
		} else {
			resource += fmt.Sprintf("/%s", r)
		}
		res = append(res, resource)
	}

	var this_handler Handler
	var result string

	for i := len(res) - 1; i >= 0; i-- {
		resource := res[i]

		this_handler, err = this.GetHandler(resource)
		if err == nil {
			result, err = this_handler.Process(r, version, resource, this_handler.ProcessFunc)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, result) //MakeErrorResult(-1, err.Error()))
			} else {
				header.Add("Content-Length", fmt.Sprintf("%v", len(result)))
				io.WriteString(w, result)
			}
			return
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, MakeErrorResult(-1, err.Error()))
	}

	return
}

func (this *Router) GetHandler(resource string) (Handler, error) {
	handler, found := this.routes[resource]
	if found && handler != nil {
		return handler, nil
	} else {
		return nil, errors.New("handler not found.")
	}
}

//
//通过正则表达式选择路由程序
//
func (this *Router) ParseURL(url string) (version int, resources []string, err error) {
	urlPattern := "/v(\\d+)/(\\w+)"
	urlRegexp, err := regexp.Compile(urlPattern)
	if err != nil {
		return
	}
	matchs := urlRegexp.FindStringSubmatch(url)
	if matchs == nil {
		err = errors.New("Wrong Request URL")
		return
	}

	/*
	   for i, str := range matchs {
	       fmt.Println(i, ": ", str)
	   }
	*/

	versionNum, _ := strconv.ParseInt(matchs[1], 10, 8)
	version = int(versionNum)

	for i := 2; i < len(matchs); i++ {
		resources = append(resources, matchs[i])
	}

	return
}

func MakeErrorResult(errcode int, errmsg string) string {
	data := map[string]interface{}{
		"error_code": errcode,
		"message":    errmsg,
	}
	result, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("{\"error_code\":%v,\"message\":\"%v\"}", errcode, errmsg)
	}
	return string(result)
}
