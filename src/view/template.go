package view

import (
	"net/http"
	"../global"
	"html/template"
	"time"
	"fmt"
)

//go模板函数
var funcMap=template.FuncMap{
	"noescape": func(s string) template.HTML{
		return template.HTML(s)
	},
	"formatTime": func(t time.Time,layout string) string {
		return t.Format(layout)
	},
}

//render 渲染模板并输出
func Render(w http.ResponseWriter,r *http.Request,html string,data map[string]interface{}){
	if data==nil{
		data=make(map[string]interface{})
	}

	data["app"]=global.App.Name
	//data["site_name"]=config.YamlConfig.Get("setting.site_name").String()
	//data["title"]=config.YamlConfig.Get("setting.title").String()
	//data["subtitle"]=config.YamlConfig.Get("setting.subtitle").String()

	//加载布局模板layout.html
	tpl, err := template.New("layout.html").Funcs(funcMap).
		ParseFiles("C:/Users/Administrator/Desktop/goExample/myDreamGo/template/theme/default/layout.html")
	if err!=nil{
		//服务器内部错误
		fmt.Println("服务器内部错误:"+err.Error())
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	//渲染模板，并输出到w
	fmt.Println("渲染模板")
	fmt.Println(data)
	err = tpl.Execute(w, data)
	if err!=nil{
		//服务器内部模板渲染错误
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}