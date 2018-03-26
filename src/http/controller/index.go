package controller

import (
	"net/http"
	"../../route"
	"../../datasource"
	"../../view"
	"fmt"
)

//通过以下方式都可以访问到首页
var defaults=map[string]bool{
	"/":true,
	"/index.html":true,
	"/index.htm":true,
}

type IndexController struct {

}
//注册首页的路由
func (self IndexController)RegistRoute(){
	route.HandleFunc("/",self.Home)
}

//首页
func (IndexController)Home(w http.ResponseWriter,r *http.Request){
	if ok:=defaults[r.RequestURI];!ok{
		//抛出首页未找到的错误
		http.NotFound(w,r)
		return
	}

	posts := datasource.DefaultDataSourcer.PostList()
	fmt.Println("文章列表的查询结果:")
	for _,post:= range posts{
		fmt.Println("文章标题:"+post.Title)
	}

	fmt.Println("渲染首页")
	view.Render(w,r,"index.html",map[string]interface{}{"post":posts})

}


