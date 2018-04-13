package main

import (
	"../global"
	"../datasource"
	"../http/controller"
	"net/http"
	"../route"
	"fmt"
	"os"
	"bufio"
)

var configData=&InitConfigYamlStruct{}

func main() {

	//读取配置文件
	LoadConfig()
	//初始化数据库
	fmt.Println("初始化数据库")
	datasource.Init()
	//从配置文件中获取监听IP和端口
	fmt.Println("监听ip以及端口")
	global.App.Host="0.0.0.0"
	global.App.Port="80"
	fmt.Println("ip地址:"+global.App.Host)
	fmt.Println("端口:"+global.App.Port)
	addr:=global.App.Host+":"+global.App.Port
	//注册路由
	fmt.Println("注册路由")
	fmt.Println("启动监听")
	controller.RegisterRoutes()
	http.ListenAndServe(addr,route.DefaultBlogMux)


}

/**
	加载配置文件
 */
func LoadConfig(){
	//打开配置文件并获取
	configDatas, err:= os.Open("./src/config/env.yml")
	defer configDatas.Close()
	//读取配置文件
	if err!=nil{
		fmt.Println("读取文件错误"+err.Error())
		return
	}
	reader := bufio.NewReader(configDatas)
	var fileStr string
	for{
		string, err := reader.ReadString('\n')
		if err!=nil && err.Error()=="EOF"{
			fmt.Println("读取文件中的字符串错误:"+err.Error())
			fmt.Println("fuck"+fileStr)
			return
		}
		if string!=""{
			fileStr=fileStr+string
		}
	}
	//err = yaml.Unmarshal([]byte(configDatas), &configData)
	if err!=nil{
		fmt.Println("转化配置文件错误")
		return
	}
	fmt.Printf("what the fuck: %v",fileStr)


}

//读取yaml的配置文件
type InitConfigYamlStruct struct {

	Listen struct{
		Host string `yaml:"host"`
		Port int64	`yaml:"post"`
	}

	Setting struct{
		Site_name string `yaml:"site_name"`
		Title string `yaml:"title"`
		SubTitle string `yaml:"subtitle"`
	}

	Seo struct{
		Keywords string `yaml:"keywords"`
		Description string `yaml:"description"`
	}

	DataSource struct{
		DatasourceType string `yaml:"type"`
		Url string `yaml:"url"`
		Monogdbaddr string `yaml:"monogdbaddr"`
		Monogdbdb string `yaml:"monogdbdb"`
		MysqlAddr string `yaml:"mysqlAddr "`
	}

	Theme string
}

