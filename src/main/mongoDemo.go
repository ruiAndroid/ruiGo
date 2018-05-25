package main
import("gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	//连接mongo数据库
	session,err:=mgo.Dial("mongodb://rui:jianrui123@120.79.186.178:27017/rui")
	defer session.Close()
	if err!=nil{
		fmt.Println("连接错误")
		fmt.Println(err)
		return
	}
	fmt.Println(session)
	//设置session的mode
	session.SetMode(mgo.Monotonic, true)
	//通过Session进行增删改查
	C:=session.DB("rui").C("bug_info")
	bug_info:=make([]BugInfo,100)
	//先插入一个试试
	err = C.Insert(&BugInfo{
		bson.NewObjectId(),
		"测试信息",
	})
	if err!=nil{
		fmt.Println("插入失败")
		return
	}
	err = C.Find(nil).All(&bug_info)
	if err!=nil{
		fmt.Println("查询所有的bug信息错误")
		return
	}
	fmt.Printf("查询的结果:%v",bug_info)

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

	//删除一条试试
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
	//}
}

type User struct {
	Id_ bson.ObjectId `bson:"_id"`
	UserName  string `bson:"user_name"`
	Age string `bson:"age"`
}

type BugInfo struct {
	Id_ bson.ObjectId `bson:"_id"`
	BugMsg string `bson:"bug_msg"`
}