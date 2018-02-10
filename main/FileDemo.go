package main

import (
	"os"
)

/**
go 语言文件处理

//创建名称为name的目录,权限是perm 如0777
func mkDir(name string,perm FileMode)error
//根据path创建多级目录
func mkDirAll(path string,perm FileMode)error
//删除名称为name的目录，当前目录下有文件或者其他目录时会报错
func remove(name string) error
//删除名称为name的所有多级目录，如果path是单个名称,那么该目录下的子目录全部删除
func removeAll(name string)error


建立和打开文件
//根据提供的文件名创建一个新文件,返回一个文件对象,默认权限是0666的文件,返回的文件对象是可读写的
func create(name string)(file *File,error Error)
//根据文件描述符创建相应的文件，返回一个文件对象
func newFile(fd uintptr,name File)*File

通过如下两个方法来打开文件：
func Open(name string) (file *File, err Error)
该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。
func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限

写文件函数：
func (file *File) Write(b []byte) (n int, err Error)
写入byte类型的信息到文件
func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
在指定位置开始写入byte类型的信息
func (file *File) WriteString(s string) (ret int, err Error)
写入string信息到文件

读文件函数：
func (file *File) Read(b []byte) (n int, err Error)
读取数据到b中
func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
从off开始读取数据到b中

 */
func main() {
	//os.Mkdir("rui",0777)
	//os.Remove("rui")
	//创建一个文件
	//result:=make([]byte,1024)
	file:="rui.txt"
	os.Remove(file)
	//open, _ := os.Open(file)
	//defer open.Close()
	//n, _ := open.Read(result)
	//fmt.Println(string(result[:n]))
	/*create, e := os.Create(file)
	if e!=nil{
		fmt.Println("create file fail")
		return
	}
	defer create.Close()
	for i:=0;i<5;i++{
		create.WriteString("aaa\n")
		create.Write([]byte("bbb\n"))
	}*/

}
