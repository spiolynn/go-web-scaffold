package main

import (
	"fmt"
	"go-web-scaffold/models"
	"go-web-scaffold/pkgs/logging"
	"go-web-scaffold/pkgs/setting"
	"go-web-scaffold/routers"
	"net/http"
	"time"
)

/*
主程序
*/

// 初始化
func _init() {
	logging.Info(" server init start ")

	logging.Info(" db init start ")
	models.OpenDB()

	logging.Info(" redis init start ")

	logging.Info(" server init end ")

}

func _start_server() {

	// 路由初始
	router := routers.InitRouter()
	// web服务初始
	s := &http.Server{
		// 监听的TCP地址
		Addr:           fmt.Sprintf(":%d", setting.G_cfg_yaml.Server.Httpprot),
		Handler:        router,
		ReadTimeout:    time.Duration(setting.G_cfg_yaml.Server.Readtimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.G_cfg_yaml.Server.Wirtetimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 服务启动
	s.ListenAndServe()

}

func main() {
	_init()
	_start_server()
}
