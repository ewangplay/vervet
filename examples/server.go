package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/ewangplay/vervet"
	"github.com/outmana/log4jzl"
)

func main() {
	var err error

    
	//读取启动参数
	var configFile string
	flag.StringVar(&configFile, "config", "config.ini", "configure file full path")
	flag.Parse()

	//读取配置文件
	config, err := NewConfigure(configFile)
	if err != nil {
		fmt.Printf("[ERROR] Parse Configure File Error: %v\n", err)
		return
	}

	//启动日志系统
	logger, err := log4jzl.New("ProxyServer")
	if err != nil {
		fmt.Printf("[ERROR] Create logger Error: %v\n", err)
		return
	}

	base_handler := vervet.NewBaseHandler(logger)
	routes := map[string]vervet.Handler{
		"demo": &DemoHandler{base_handler},
        "demo/animal": &DemoAnimalHandler{base_handler},
	}
	router := vervet.NewRouter(config, routes)

	addr := ":8089"
	logger.Info("Server Start..., Listening on %v", addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		logger.Error("Server start fail: %v", err)
		os.Exit(1)
	}
}
