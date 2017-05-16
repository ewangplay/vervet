package vervet

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type BaseHandler struct {
	Logger Logger
}

func NewBaseHandler(logger Logger) *BaseHandler {
    return &BaseHandler{logger}
}

func (this *BaseHandler) Process(r *http.Request, resources []string, f ProcessFunc) (string, error) {
	var err error
	var body []byte
	var startTime time.Time

	startTime = time.Now()

	result := make(map[string]interface{})

    //Generate the log id
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	log_id1 := rand.Intn(100000)
	log_id2 := rand.Intn(100000)
	log_id := fmt.Sprintf("%d%d", log_id1, log_id2)
	result["log_id"] = log_id

	this.Logger.Info("[LOG_ID:%v] [METHOD:%v] [URL:%v]", log_id, r.Method, r.RequestURI)

    //Parse the request parameters
	params, err := this.parseArgs(r)
	if err != nil {
		result["error_code"] = -1
		result["message"] = "parse request parameter error" //err.Error()
		goto END
	}

    //Read the request body
	body, err = ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		result["error_code"] = -1
		result["message"] = "read request body error" //err.Error()
		goto END
	}

    //Perform the actual business process
	err = f(r.Method, resources, params, body, result)
	if err != nil {
		result["error_code"] = -1
		if strings.HasPrefix(err.Error(), "[ERROR_INFO]") {
			result["message"] = strings.TrimPrefix(err.Error(), "[ERROR_INFO]")
		} else {
			result["message"] = fmt.Sprintf("systerm error! LOG_ID: %v", log_id)
		}
		goto END
	}

	result["error_code"] = 0

END:
	if err != nil {
		this.Logger.Error("[LOG_ID:%v] %v", log_id, err)
		if string(body) != "" {
			this.Logger.Error("[LOG_ID:%v] [Request Body : %v]", log_id, string(body))
		}
		this.Logger.Error("[LOG_ID:%v] [Response Result : %v]", log_id, result)
	}

	result["cost"] = fmt.Sprintf("%v", time.Since(startTime))
	result["request_url"] = r.RequestURI

	this.Logger.Info("[LOG_ID:%v] [COST:%v]", log_id, result["cost"])
	resStr, _ := this.createJSON(result)

	return resStr, err
}

func (this *BaseHandler) createJSON(result map[string]interface{}) (string, error) {
	r, err := json.Marshal(result)
	if err != nil {
		this.Logger.Error("%v", err)
		return "", err
	}
	return string(r), nil
}

func (this *BaseHandler) parseArgs(r *http.Request) (map[string]string, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	//每次都重新生成一个新的map，否则之前请求的参数会保留其中
	res := make(map[string]string)
	for k, v := range r.Form {
		res[k] = v[0]
	}

	return res, nil
}
