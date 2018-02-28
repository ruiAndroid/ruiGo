package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
gin 框架demo
 */
func main() {
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