package main

import (
	"net/http"
	"fmt"
)
var urls=[]string{
	"http://www.baidu.com",
	"http://www.taobao.com",
}


/**
http.Get获取网页内容
get返回值的res包含了网页内容

 */
const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

func main() {

	//resp, _ := http.Get("http://www.baidu.com")
	//reader, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%q",string(reader))

	//for _,url:=range urls{
	//	resp, _ := http.Head(url)
	//	fmt.Println(url,":"+resp.Status)
	//}

	//http.HandleFunc("/",HelloServer)
	//http.ListenAndServe("localhost:8080",nil)

	//一个简单的网页应用

}

func HelloServer(w http.ResponseWriter,request *http.Request){
	
	fmt.Println("inSide hellow handler")
	fmt.Fprintf(w,"<h1>%s<h1><div>%s</div>", "asdsadasdasdasd", request.URL.Path[1:])


}

