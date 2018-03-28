package controller

import (
	"net/http"
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

type HttpJson struct {
	Msg string `json:"msg"`
	Code string `json:"code"`
}

//注册用户bug收集的路由
func (my UserBugTrackerController)RegistRoute(){
	http.HandleFunc("/user/",my.InsertUserBug)
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
	fmt.Println("请求的方式:"+r.Method)
	/*if r.Method=="GET"{
		fmt.Println("GET请求")
		fmt.Println("userData:", r.Form["userData"])
		bytes, _ := ioutil.ReadAll(r.Body)
		fmt.Println("fuckasdasdas:"+string(bytes))
	}else{
		fmt.Println("POST请求")

	}*/
	var value string
	httpJson:=&HttpJson{}
	if len(r.Form["userData"]) > 0 && r.Form["userData"][0]!=""{
		value=r.Form["userData"][0]
		i := len(value)
		fmt.Println(i)
		//转化json
		userBugTrackInfo :=&datasource.UserBugTrack{}
		err := json.Unmarshal([]byte(value), userBugTrackInfo)
		if err!=nil{
			httpJson=&HttpJson{Msg:"error",Code:"500"}
			json.NewEncoder(w).Encode(httpJson)
			fmt.Fprintln(w, httpJson)
			return
		}
		//插入数据库
		fmt.Printf("客户端传过来的json: %v",userBugTrackInfo)
		result:=datasource.InsertUserBugInfo(userBugTrackInfo)
		fmt.Printf("%b",result)
		if result{
			httpJson=&HttpJson{Msg:"ok",Code:"200"}
			json.NewEncoder(w).Encode(httpJson)
			fmt.Fprintln(w, httpJson)
		}else{
			httpJson=&HttpJson{Msg:"error",Code:"500"}
			json.NewEncoder(w).Encode(httpJson)
			fmt.Fprintln(w, httpJson)
		}
		//fmt.Println(httpJson)
	}else{
		fmt.Println("error")
		//返回Json告诉客户端格式有问题
		httpJson=&HttpJson{Msg:"error",Code:"500"}
		json.NewEncoder(w).Encode(httpJson)
		fmt.Fprint(w, httpJson)
	}

}
