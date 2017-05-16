package main

import (
	"errors"
	"github.com/ewangplay/config"
	"strconv"
)

type Configure struct {
	ConfigureMap map[string]string
}

func NewConfigure(filename string) (*Configure, error) {
	config := &Configure{}

	config.ConfigureMap = make(map[string]string)
	err := config.ParseConfigure(filename)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (this *Configure) loopConfigure(sectionName string, cfg *config.Config) error {

	if cfg.HasSection(sectionName) {
		section, err := cfg.SectionOptions(sectionName)
		if err == nil {
			for _, v := range section {
				options, err := cfg.String(sectionName, v)
				if err == nil {
					this.ConfigureMap[v] = options
				}
			}

			return nil
		}
		return errors.New("Parse Error")
	}

	return errors.New("No Section")
}

func (this *Configure) ParseConfigure(filename string) error {
	cfg, err := config.ReadDefault(filename)
	if err != nil {
		return err
	}

	this.loopConfigure("sever", cfg)
	this.loopConfigure("service", cfg)

	return nil
}

func (this *Configure) GetPort() (int, error) {

	portstr, ok := this.ConfigureMap["port"]
	if ok == false {
		return 9090, errors.New("No Port set, use default")
	}

	port, err := strconv.Atoi(portstr)
	if err != nil {
		return 9090, err
	}

	return port, nil
}

func (this *Configure) GetUrlPattern() (string, error) {

	UrlPattern, ok := this.ConfigureMap["urlpattern"]
	if ok == false || UrlPattern == "" {
		return "", errors.New("No UrlPattern Setting")
	}

	return UrlPattern, nil
}

func (this *Configure) GetLogIdLiteral() (string, error) {
    logIdLiteral, ok := this.ConfigureMap["log_id_literal"]
    if ok == false || logIdLiteral == "" {
        return "", errors.New("No log id literal setting")
    }
    return logIdLiteral, nil
}

func (this *Configure)   GetErrorCodeLiteral() (string, error) {
    errorCodeLiteral, ok := this.ConfigureMap["error_code_literal"]
    if ok == false || errorCodeLiteral == "" {
        return "", errors.New("No error code literal setting")
    }
    return errorCodeLiteral, nil
}

func (this *Configure)   GetErrorMessageLiteral() (string, error) {
    errorMsgLiteral, ok := this.ConfigureMap["error_message_literal"]
    if ok == false || errorMsgLiteral == "" {
        return "", errors.New("No error message literal setting")
    }
    return errorMsgLiteral, nil

}

func (this *Configure)   GetTimeCostLiteral() () (string, error) {
    timeCostLiteral, ok := this.ConfigureMap["time_cost_literal"]
    if ok == false || timeCostLiteral == "" {
        return "", errors.New("No time cost literal setting")
    }
    return timeCostLiteral, nil

}

func (this *Configure)   GetRequestUrlLiteral() (string, error) {
    requestUrlLiteral, ok := this.ConfigureMap["request_url_literal"]
    if ok == false || requestUrlLiteral == "" {
        return "", errors.New("No request url literal setting")
    }
    return requestUrlLiteral, nil

}


