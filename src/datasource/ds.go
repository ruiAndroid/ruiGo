package datasource

import (
	"../model"
	"../config"
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
	DefaultDataSourcer = NewMySql(config.YamlConfig.Get("datasource.mysqlAddr").String())
	go DefaultDataSourcer.UpdateDataSource()
}
