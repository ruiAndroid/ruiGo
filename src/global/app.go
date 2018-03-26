package global

import (
	"time"
	"sync"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"fmt"
	"io"
)

//global 全局信息


//Build构建信息,从git 仓库获取
var Build string

type app struct {
	//名称
	Name string
	//构建信息
	Build string
	//版本号
	Version string
	//构建时间
	BuildDate time.Time

	//项目目录
	ProjectRoot string
	//模板目录
	TemplateDir string

	//版权信息
	Copyright string
	LaunchTime time.Time

	//主机
	Host string
	//端口
	Port string

	locker sync.Mutex
}

//App is the App Info
var App =&app{

}

var showVersion=flag.Bool("version",false,"print version of this binary ")

/*
 初始化
 */
func init(){
	App.Name=os.Args[0]
	App.Version="V1.0.0"
	App.Build=Build
	//初始化编译时间为当前时间
	App.LaunchTime=time.Now()
	//查找可执行程序的路径
	binaryPath,err:=exec.LookPath(os.Args[0])
	if err!=nil{
		panic(err)
	}
	//获取可执行程序的绝对路径
	binaryPath,err=filepath.Abs(binaryPath)
	if err!=nil{
		panic(err)
	}
	//获取可执程序的文件信息
	fileInfo ,err :=os.Stat(binaryPath)
	if err!=nil{
		panic(err)
	}
	//构建时间为可执行程序的修改时间
	App.BuildDate=fileInfo.ModTime()
	App.Copyright=fmt.Sprintf("%d",time.Now().Year())

}

//初始化路径
func (this *app)InitPath(){
	//设置项目根目录路径
	App.setProjectRoot()
	//设置模板的存放路径
	App.setTemplateDir("default")


}


//设置项目根目录
func (this *app)setProjectRoot(){
	currFileName:=os.Args[0]
	binaryPath,err:=exec.LookPath(currFileName)
	if err!=nil{
		panic(err)
	}
	binaryPath,err =filepath.Abs(binaryPath)
	if err!=nil{
		panic(err)
	}
	projectRoot:=filepath.Dir(filepath.Dir(binaryPath))
	this.ProjectRoot=projectRoot+"/"
}


//设置模板目录
func (this *app)setTemplateDir(theme string){
	this.TemplateDir=this.ProjectRoot+"template/theme/"+theme

}

func PrintVersion(writer io.Writer){
	if !flag.Parsed(){
		flag.Parse()
	}

	if showVersion==nil || !*showVersion{
		return
	}

	fmt.Fprintf(writer,"Binary:%s\n",App.Name)
	fmt.Fprintf(writer,"Version:%s\n",App.Version)
	fmt.Fprintf(writer,"Build:%s\n",App.Build)
	fmt.Fprintf(writer,"Complete date:%s\n",App.BuildDate.Format("2006-01-02 15:04:05"))
	os.Exit(0)

}











