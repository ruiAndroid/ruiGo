package utils

import (
	"os"
	"strings"
)

//Exist检查文件或目录是否存在
func Exist(filename string)bool{
	_, err := os.Stat(filename)
	return err==nil ||os.IsExist(err)
}

//ScanDir列出指定路径中的文件和目录
//如果目录不存在,则返回空slice
func ScanDir(directory string)[]string{
	file,err:=os.Open(directory)
	if err!=nil{
		return []string{}
	}
	names, err := file.Readdirnames(-1)
	if err!=nil{
		return []string{}
	}
	return names
}

//判断给定文件名是否为一个正常的文件
//如果文件存在且为正常的文件则返回true
func IsDir(filename string)bool{
	return isDirOrFile(filename,true)
}

//判断给定的文件名是否为一个正常的文件
func IsFile(filename string)bool{
	return isDirOrFile(filename,false)
}

func Filename(file string)string{
	if file==""{
		return ""
	}
	//返回.在文件名中的位置
	index := strings.LastIndex(file, ".")
	//返回去掉.拓展名的完整文件名
	return file[:index]
}

//判断是文件还是目录,根据decideDir为true表示判断是否为目录,否则判断是否为文件
func isDirOrFile(filename string,decideDir bool)bool{
	fileInfo, err := os.Stat(filename)
	if err!=nil{
		return false
	}
	//判断是文件还是目录
	isDir := fileInfo.IsDir()
	if decideDir{
		return isDir
	}
	return !isDir
}