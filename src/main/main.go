package main

import (
	"../global"
	"../datasource"
	"../http/controller"
	"net/http"
	"../route"
	"fmt"
)

func main() {
	//初始化数据库
	fmt.Println("初始化数据库")
	datasource.Init()
	//从配置文件中获取监听IP和端口
	fmt.Println("监听ip以及端口")
	global.App.Host="0.0.0.0"
	global.App.Port="80"
	fmt.Println("ip地址:"+global.App.Host)
	fmt.Println("端口:"+global.App.Port)
	addr:=global.App.Host+":"+global.App.Port
	//注册路由
	fmt.Println("注册路由")
	fmt.Println("启动监听")
	controller.RegisterRoutes()
	http.ListenAndServe(addr,route.DefaultBlogMux)

}
