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

	resource, version, err := this.ParseURL(r.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, MakeErrorResult(-1, err.Error()))
	} else {
		this_handler, err := this.GetHandler(resource)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, MakeErrorResult(-1, err.Error()))
		} else {
			result, err := this_handler.Process(r, version, resource, this_handler.ProcessFunc)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, result) //MakeErrorResult(-1, err.Error()))
			} else {
				header.Add("Content-Length", fmt.Sprintf("%v", len(result)))
				io.WriteString(w, result)
			}
		}
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
func (this *Router) ParseURL(url string) (resource string, version int, err error) {
	//确定是否是本服务能提供的控制类型
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
	resource = matchs[2]

    //TODO: 把URL中的多级资源保存

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
