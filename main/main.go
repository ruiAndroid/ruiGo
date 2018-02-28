package main

import (
	"net/http"
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
)

/**
部署go应用
Supervisord
Supervisord会帮你把管理的应用程序转化为daemon程序,并且可以通过命令执行开启，关闭，重启等操作
并且崩溃后会自动重启,保证程序有自我修复功能
*/
func hh(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "然山海不可平啊") //这个写入到w的是输出到客户端的
}

func main() {
	/*http.HandleFunc("/", hh) //设置访问的路由
	http.ListenAndServe(":80", nil) //设置监听的端口*/
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	
	router := gin.Default()
	/*
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"gin demo is running")
	})
	*/
	//restful路由
	/*router.GET("/:user", func(context *gin.Context) {
		user := context.Param("user")
		context.String(http.StatusOK,"您好啊 %s",user)
	})*/
	//参数处理
	router.GET("/user", func(context *gin.Context) {
		name:=context.Query("name")
		context.String(http.StatusOK,"fuck the %s",name)
	})
	/*	//body处理
		router.GET("/user", func(context *gin.Context) {

		})*/
	router.Run(":80")
}