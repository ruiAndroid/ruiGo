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
	"gopkg.in/yaml.v2"
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
	//从配置文件中读取信息
	fmt.Println("端口号:"+configData.Listen.Port)
	fmt.Println("ip地址:"+configData.Listen.Host)
	if (configData.Listen.Host!="")&& (configData.Listen.Port!="") {
		global.App.Host=configData.Listen.Host
		global.App.Port=configData.Listen.Port
		addr:=global.App.Host+":"+global.App.Port
		//注册路由
		fmt.Println("注册路由")
		fmt.Println("启动监听")
		controller.RegisterRoutes()
		http.ListenAndServe(addr,route.DefaultBlogMux)
	}else{
		fmt.Println("读取配置文件错误")
		return
	}


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
			//fmt.Println("fuck"+fileStr)
			break
		}
		if string!=""{
			fileStr=fileStr+string
		}
	}

	err = yaml.Unmarshal([]byte(fileStr), &configData)
	if err!=nil{
		fmt.Println("转化配置文件错误")
		return
	}else{
		fmt.Printf("配置文件信息: %v",configData)
	}
}

//读取yaml的配置文件
type InitConfigYamlStruct struct {

	Listen struct{
		Host string `yaml:"host"`
		Port string	`yaml:"port"`
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

