package main
import(
	"fmt"
	"bufio"
	"gopkg.in/yaml.v2"
	"os"
	"net/http"
	"../global"
	"../http/controller"
	"../route"
	"../datasource"
)

var config=&InitConfigStruct{}

func main() {
	//加载配置文件
	loadSysConfig()
	datasource.Init()
	if (config.Listen.Host!="")&& (config.Listen.Port!="") {
		global.App.Host=config.Listen.Host
		global.App.Port=config.Listen.Port
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
	//fmt.Println(session)
	////设置session的mode
	//session.SetMode(mgo.Monotonic, true)
	////通过Session进行增删改查
	//C:=session.DB("rui").C("bug_info")
	//bug_info:=make([]BugInfo,100)
	//
	////先插入一个试试
	//err = C.Insert(&BugInfo{
	//	bson.NewObjectId(),
	//	"测试apkVersion",
	//	"测试SysVersion",
	//	"测试PhoneModel",
	//	"测试UserPhone",
	//	"测试UserName",
	//	time.Now().String(),
	//	"测试bugMsg",
	//})
	//
	//if err!=nil{
	//	fmt.Println("插入失败")
	//	return
	//}
	//err = C.Find(nil).All(&bug_info)
	//if err!=nil{
	//	fmt.Println("查询所有的bug信息错误")
	//	return
	//}
	//fmt.Printf("查询的结果:%v",bug_info)

	//创建切片用来存储从数据库中查询到的数据
	//user:=make([]User,100)
	////查询所有的user表中的信息
	//err=C.Find(nil).All(&user)
	//if err!=nil{
	//	fmt.Println("查询错误:"+err.Error())
	//	return
	//}
	//fmt.Printf("查询的结果 %v:",user)
	//然后插入一条试试？
	/*err= C.Insert(&User{
		bson.NewObjectId(),
		"ruirui",
		"22",
	})
	if(err!=nil){
		fmt.Println("插入错误:"+err.Error())
		return
	}*/

	//err = C.Remove(bson.M{"user_name": "ruirui"})
	//if err!=nil{
	//	fmt.Println("删除错误:"+err.Error())
	//	return
	//}
	//更新一条试试
	//err = C.Update(bson.M{"_id": bson.ObjectIdHex("5b00014925544e16839171e9")}, bson.M{"$set": bson.M{"user_name": "rui"}})
	//if err!=nil{
	//	fmt.Println("更新错误:"+err.Error())
	//}
	//字段增加值试试
	//err = C.Update(bson.M{"_id": bson.ObjectIdHex("5b00014925544e16839171e9")}, bson.M{"$inc": bson.M{"gender":"man",}})
	//if err!=nil{
	//	fmt.Println("字段增加值错误:"+err.Error())
	//	return
	//删除一条试试
	//}

}
/**
加载系统配置文件
 */
func loadSysConfig(){
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

	err = yaml.Unmarshal([]byte(fileStr), &config)
	if err!=nil{
		fmt.Println("转化配置文件错误")
		return
	}else{
		fmt.Printf("配置文件信息: %v",config)
	}

}



//读取yaml的配置文件
type InitConfigStruct struct {

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
