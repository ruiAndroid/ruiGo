package controller

import (
	"net/http"
	"../../route"
	"fmt"
	"encoding/json"
	"../../datasource"

)

//通过以下方式都可以访问到首页
/*var userBugDefaults=map[string]bool{
	"/":true,
	"/user/userBugTrack.html":true,
	"/user/userBugTrack.htm":true,
}*/

type UserBugTrackerController struct {

}

//注册用户bug收集的路由
func (my UserBugTrackerController)RegistRoute(){
	route.HandleFunc("/user/",my.InsertUserBug)
}

//收集用户bug的接口
func (UserBugTrackerController)InsertUserBug(w http.ResponseWriter,r *http.Request){
/*	if ok:=userBugDefaults[r.RequestURI];!ok{
		//抛出首页未找到的错误
		fmt.Println("未找到该网页")
		http.NotFound(w,r)
		return
	}*/
	r.ParseForm()
	var value string
	if len(r.Form["userData"]) > 0 && r.Form["userData"][0]!=""{
		value=r.Form["userData"][0]
		//转化json
		userBugTrackInfo :=&datasource.UserBugTrack{}
		err := json.Unmarshal([]byte(value), userBugTrackInfo)
		if err!=nil{
			fmt.Fprintln(w, "json格式有问题兄弟")
			return
		}
		//插入数据库
		fmt.Printf("json: %v",userBugTrackInfo)
		result:=datasource.InsertUserBugInfo(userBugTrackInfo)
		fmt.Printf("%b",result)
		fmt.Fprintln(w, value)
	}else{
		//返回Json告诉客户端格式有问题
		fmt.Fprintln(w, "格式有问题兄弟")
	}

}
