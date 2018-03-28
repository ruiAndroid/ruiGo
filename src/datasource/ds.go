package datasource

import (
	"../model"
)
const (
	TypeGir="git"
	TypeMysql="mysql"
)

type DataSourcer interface {
	PostList()[]*model.Post
	UpdateDataSource()

}

// DefaultDataSourcer 默认数据源
var DefaultDataSourcer DataSourcer

//Init数据源初始化
func Init(){
	DefaultDataSourcer = NewMySql("root:jianrui123@tcp(localhost:3306)/rui")
	go DefaultDataSourcer.UpdateDataSource()
}
