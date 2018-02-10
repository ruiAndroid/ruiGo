package main

import (
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Note struct {
	Id_ string  `bson:"_id"`
	WordId     int64 `bson:"word_id"`
	UserId     string `bson:"user_id"`
	Author     string `bson:"author"`
	Content    string `bson:"content"`
	Timestamp  string `bson:"timestamp"`
	Like       int64 `bson:"like"`
	Reported   int64 `bson:"reported"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
}
func main() {
	dialInfo:=&mgo.DialInfo{
		Addrs:     []string{"139.224.47.75:27017"},
		Direct:    false,
		Timeout:   time.Second * 10,
		Database:  "lion-test",
		Username:  "lion-test",
		Password:  "Test-EC.98",
	}

	session, e := mgo.DialWithInfo(dialInfo)
	if e!=nil{
		fmt.Println("链接错误",e.Error())
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//notes:=[]Note{}
	//note:=Note{}
	collection := session.DB("lion-test").C("note")

	//collection.Find(nil).All(&notes)
	//objectId:=bson.ObjectIdHex("58f9f5b9362d0e2cf0225b51")
	//collection.FindId(objectId).One(&note)

	//$or 查询
	//collection.Find(bson.M{"$or":[]bson.M{bson.M{"like":bson.M{"$eq":16}},bson.M{"like":bson.M{"$eq":100}},}}).All(&notes)
	//fmt.Printf("%+v \n",notes)

	//修改
	collection.Update(bson.M{"_id":bson.ObjectIdHex("58fa01a4362d0e31fc1d0f91")},
	bson.M{"$pull":bson.M{
		"test":100}})
}
