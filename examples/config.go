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

//服务信息
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

//服务控制信息
func (this *Configure) GetUrlPattern() (string, error) {

	UrlPattern, ok := this.ConfigureMap["urlpattern"]
	if ok == false || UrlPattern == "" {
		return "/v(\\d)/(contents|control|hello|email|sms|template)/", errors.New("No UrlPattern,use defualt")
	}

	return UrlPattern, nil
}

