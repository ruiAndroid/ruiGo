package controller

import (
	"../../route"
	"net/http"
)
type PostController struct{

}

//注册路由
func (self PostController)RegistRoute(){
	route.HandleFunc("/post/",self.Detail)
}

//处理文件详情请求
func (PostController)Detail(w http.ResponseWriter,r *http.Request){
	//获取文章文件名,即文章的路径
	//filename := filepath.Base(r.RequestURI)


}
