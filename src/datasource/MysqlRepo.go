package datasource

import (
	"fmt"
	"io/ioutil"
	"../global"
	"gopkg.in/yaml.v2"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"../utils"
	"os"
	"../model"
	"bytes"
)

const(
	//PostDir文章存放目录
	PostDir="data/post/"

	//IndexFile首页数据文件
	IndexFile="index.yaml"

	//Archive 归档数据文件
	ArchiveFile="archive.yaml"

	//TagFile标签文件
	TagFile="tags.yaml"

	//FriendFile 友情链接数据文件
	FriendFile="friends.yaml"
)

//MysqlRepo mysql 数据源结构体
type MysqlRepo struct {
	db						*sql.DB
	selectTag				*sql.Stmt
	selectArticleById 		*sql.Stmt
	selectArticleIndex		*sql.Stmt
	selectArticleTagsById 	*sql.Stmt
	selectArticleArchives	*sql.Stmt
	selectArticlesByTag		*sql.Stmt
	selectFriends			*sql.Stmt
	//插入用户bug信息
	insertNewUserBug      	*sql.Stmt
}

//文章的详情Struts
type articleInfo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	PubTime string `json:"pub_time"`
	Content string `json:"content"`
	Path string `json:"path"`
	PostTime string `json:"post_time"`
	Tags string `json:"tags"`
}
/**
收集用户bug的struct
 */
type UserBugTrack struct {
	UserId string `json:"user_id"`
	DownloadTime string `json:"download_time"`
	PhoneModel string `json:"phone_model"`
	SdCardMemory string `json:"sd_card_memory"`
	OriginalZipSize string `json:"original_zip_size"`
	WordId []string `json:"word_id"`
	AudioFileCount string `json:"audio_file_count"`
	PicFileCount string `json:"pic_file_count"`
	BugWordId string `json:"bug_word_id"`
	TestOriginSize string `json:"test_origin_size"`
	TestWordId []string `json:"test_word_id"`
	TestAudioFileCount string `json:"test_audio_file_count"`
	TestPicFileCount string `json:"test_pic_file_count"`
	TestBugDate string `json:"test_bug_date"`
}

type TagInfo struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

type FriendInfo struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
	Logo string	`json:"logo"`
}
var db *sql.DB
func NewMySql(dbParams string) *MysqlRepo{
	var e error
	db, e= sql.Open("mysql", dbParams)
	if e!=nil{
		fmt.Printf("无法连接数据库"+e.Error())
	}
	return &MysqlRepo{
		db:					db,
		//selectTag: prepare(db,"SELECT * FROM `tag`"),
		selectArticleById:prepare(db,"SELECT * FROM `article_info` WHERE `id`=?"),
		selectArticleIndex:prepare(db,"SELECT * FROM `article_info` ORDER BY `pub_time` DESC LIMIT 20"),
		//selectArticleTagsById:prepare(db,"SELECT t.`name` FROM `article_tag` at LEFT_JOIN `tag` t ON at.`tag_id`=t.`id` WHERE `article_id`=?"),
		//selectArticleArchives:prepare(db,"SELECT `id`,`title`,`pub_time` FROM `article`"),
		//selectArticlesByTag:prepare(db,"SELECT a.`id`,a.`title`,a.`pub_time` FROM `article` a FROM `article` a LEFT_JOIN `article_tag` at ON a.`id`=at.`article_id` WHERE at.`tag_id`=?"),
		//查询所有的好友连接
		//selectFriends:prepare(db,"SELECT * FROM `friend_link`"),
	}
}

//数据库准备
func prepare(db *sql .DB,sql string) *sql.Stmt {
	stmt, e := db.Prepare(sql)
	if e!=nil{
		fmt.Println("数据库准备的sql出错 %s",e)
	}
	return stmt
}

//插入用户bug信息
func InsertUserBugInfo(userBugTack *UserBugTrack)bool{
	//执行插入或者更新的语句
	//先查询数据库中是否有该user_id
	/*result, e := db.Exec("SELECT * FROM user_bug_track WHERE user_id=?", userBugTack.UserId)
	if e!=nil{
		fmt.Println("向数据库插入数据错误:"+e.Error())
		return false
	}*/
	//返回插入或者更新的结果
	return true
}

func arrayToStr(strs []string) string{
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	for i:=0;i<len(strs);i++{
		if i!=len(strs)-1{
			buffer.WriteString(strs[i]+",")
		}else{
			buffer.WriteString(strs[i])
		}
	}
	fmt.Println("要传递的数组转成字符串之后的数据:"+buffer.String())
	return buffer.String()
}

//读取文章列表
func (self *MysqlRepo)PostList()[]*model.Post{
	in, e := ioutil.ReadFile(global.App.ProjectRoot + PostDir + IndexFile)
	if e!=nil{
		return nil
	}
	posts:=make([]*model.Post,0)
	err:=yaml.Unmarshal(in,&posts)
	if err!=nil {
		return nil
	}
	return posts

}

//生成首页数据文件index.yaml
func (self *MysqlRepo)GenIndexYaml(){
	//首页最多显示20篇文章
	var posts []*model.Post
	rows, err := self.selectArticleIndex.Query()
	if err!=nil{
		fmt.Println("查询首页文章列表数据错误,"+err.Error())
	}
	for rows.Next(){
		info:=articleInfo{}
		err= rows.Scan(&info.Id, &info.Title, &info.PubTime,&info.Content,
			&info.Path,&info.PostTime,&info.Tags,)
		if err!=nil{
			fmt.Println("扫描出错:"+err.Error())
		}
		posts=append(posts,self.genOnePost(info))
	}
	buf, err := yaml.Marshal(posts)
	if err != nil {
		fmt.Printf("生成index yaml error: %v\n", err)
		return
	}
	indexYaml := "C:/Users/Administrator/Desktop/goExample/myDreamGo/data/post/"+ IndexFile
	ioutil.WriteFile(indexYaml, buf, 0777)
}

//genOnePost 组装一个post
func (self *MysqlRepo)genOnePost(info articleInfo)(*model.Post){
	return &model.Post{
			Content:info.Content,
			Title:info.Title,
			Path:fmt.Sprintf("%d.html", info.Id),
			PubTime:info.PubTime,
			PostTime:info.PostTime,
			Tags:info.Tags,
		}
}

//更新数据源
func (self *MysqlRepo)UpdateDataSource(){
	//检查文章目录(data/post)是否存在,不存在则连接mysql生成
	mysqlRepoDir:="C:/Users/Administrator/Desktop/goExample/myDreamGo/data/post"
	if !utils.Exist(mysqlRepoDir){
		if err := os.MkdirAll(mysqlRepoDir, os.ModePerm);err!=nil{
			panic(err)
		}
	}

	fmt.Println("data/post的存放目录位于:"+mysqlRepoDir)
	//解析仓库文件，生成首页,归档,标签数据
	self.GenIndexYaml()
}


