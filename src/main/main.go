package main

import (
	"../global"
	"math/rand"
	"time"
	"flag"
	"../datasource"
	"../http/controller"
	"net/http"
	"../route"
	"fmt"
)
var configFile string

func Init(){
	rand.Seed(time.Now().Unix())
	flag.StringVar(&configFile,"config","C:/Users/Administrator/Desktop/goExample/myDreamGo/src/config/env.yml","The config file. Default is $ProjectRoot/config/env.yml")
	fmt.Printf(configFile)
}

func main() {
	//日志系统
	//logger:=logger.Init("dreamGo")
	//logger.Info("main...")
	//解析命令行参数
	//flag.Parse()
	//初始化程序路径
	global.App.InitPath()
	//以/开头为绝对路径,直接解析
	//Init()
	//config.Parse(configFile)
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
	controller.RegisterRoutes()
	//启动监听
	fmt.Println("启动监听")
	http.ListenAndServe(addr,route.DefaultBlogMux)


}
