package route

import (
	"net/http"
	"context"
	"time"
)


func HandleFunc(pattern string,handler func(w http.ResponseWriter,r *http.Request)){
	DefaultBlogMux.HandleFunc(pattern,handler)
}

//BlogMux路由器 拓展http.ServeMux
type BlogMux struct {
	*http.ServeMux
}
//DefaultBlogMux默认路由处理器
var DefaultBlogMux=NewBlogMux()

func NewBlogMux() *BlogMux{
	return &BlogMux{
		ServeMux:http.DefaultServeMux,
	}
}

//serverHttp路由分发方法 封装http.DefaultServeMux.ServeHttp()
func (self *BlogMux)ServeHttp(w http.ResponseWriter,r *http.Request){
	//创建上下文,并写入start_time
	ctx := context.WithValue(r.Context(), "start_time", time.Now())
	//使用上下文
	r = r.WithContext(ctx)
	//调用http.DefaultServeMux的路由分发方法
	self.ServeMux.ServeHTTP(w,r)
}

