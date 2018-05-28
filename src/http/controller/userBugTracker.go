package controller

import (
	"net/http"
	"fmt"
	"encoding/json"
	"../../datasource"
	"time"
	"strconv"

	"gopkg.in/mgo.v2/bson"
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
func (bugMsg UserBugTrackerController)RegisterBugMsgCollectionWithMgo(){
	http.HandleFunc("/bugMsg/",bugMsg.CollectionBugWithMgo)
}
//收集用户bug的接口
func (UserBugTrackerController)InsertUserBug(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	var value string
	httpJson:=&HttpJson{}

	if len(r.Form["userData"]) > 0 && r.Form["userData"][0]!=""{
		value=r.Form["userData"][0]
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
		fmt.Printf("客户端传过来的json: %v \n",userBugTrackInfo)
		result:=datasource.InsertUserBugInfo(userBugTrackInfo)
		//记录请求
		requestRecord:=&datasource.RequestRecordStruct{}
		requestRecord.UserId=userBugTrackInfo.UserId
		year:=time.Now().Year()
		month:=time.Now().Month().String()//time.Now().Month().String()
		day:=time.Now().Day()
		hour:=time.Now().Hour()
		minute:=time.Now().Minute()
		second:=time.Now().Second()

		requestRecord.RequestTime=strconv.Itoa(year)+"/"+month+"/"+strconv.Itoa(day)+"--"+strconv.Itoa(hour)+":"+strconv.Itoa(minute)+":"+strconv.Itoa(second)
		requestRecord.RequestInterface="bug track"
		datasource.RequestRecord(requestRecord)

		fmt.Printf("查询的结果:%b",result)
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


//处理请求
func (UserBugTrackerController)CollectionBugWithMgo(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	var value string
	httpJson:=&HttpJson{}
	//转化json
	bugInfoWithMongo :=&datasource.BugInfoWithMongo{}

	if len(r.Form["bugInfo"]) > 0 && r.Form["bugInfo"][0]!="" {
		value = r.Form["bugInfo"][0]
		//转化json
		err := json.Unmarshal([]byte(value), bugInfoWithMongo)
		if err != nil {
			httpJson = &HttpJson{Msg: "error", Code: "500"}
			json.NewEncoder(w).Encode(httpJson)
			fmt.Fprintln(w, httpJson)
			return
		}
		//插入数据库
		fmt.Printf("客户端传过来的json: %v \n", bugInfoWithMongo)
		bugInfoWithMongo.Id_=bson.NewObjectId()
		bugInfoWithMongo.SendTime=time.Now().String()
		//然后开始进行bug记录
		//mgo大神出马
		datasource.AddInfoToMgo(bugInfoWithMongo)


	}else{
		fmt.Println("error")
		//返回Json告诉客户端格式有问题
		httpJson=&HttpJson{Msg:"error",Code:"500"}
		json.NewEncoder(w).Encode(httpJson)
		fmt.Fprint(w, httpJson)
	}

}