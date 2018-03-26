package logger

import (
	"go.uber.org/zap"
	"../global"
	"path"
	"go.uber.org/zap/zapcore"
	"os"
	"log"
	"gopkg.in/natefinch/lumberjack.v2"
	"bytes"
)

type lumberJackWriteSyncer struct {
	*lumberjack.Logger
	buff *bytes.Buffer
	logChan chan []byte
	closeChan chan interface{}
	maxSize int
}

func (l *lumberJackWriteSyncer)run(){


}

var instance *zap.Logger

// logger instance 唯一实例
func Instance() *zap.Logger{
	return instance
}

//init 作用初始化 ,srvName 生成的日志文件夹名字
func Init(srvName string) *zap.Logger{
	instance=NewLogger(srvName)
	return instance
}

//新建日志
func NewLogger(srvName string) *zap.Logger{
	//先获取到项目的根目录
	directory:=global.App.ProjectRoot
	if len(directory)==0{	//如果根目录不存在
		directory=path.Join("loh",srvName)
	}else{
		directory=path.Join(directory,"log",srvName)
	}
	writers:=[]zapcore.WriteSyncer{newRollingFile(directory)}
	writers=append(writers,os.Stdout)

	return instance
}
/*

func newRollingFile(directory string)zapcore.WriteSyncer{

	if err:=os.MkdirAll(directory,0766);err!=nil{
		log.Println("failed create log directory:",directory,":",err)
		return nil
	}

	return newLumberJackWriteSyncer(&lumberjack.Logger{
		Filename:path.Join(directory,"output.log"),
		MaxSize:100,
		MaxAge:7,
		LocalTime:true
		Compress:false,
	})

}
*/

func newLumberJackWriteSyncer(l *lumberjack.Logger) *lumberJackWriteSyncer{
	ws:=&lumberJackWriteSyncer{
		Logger:l,
		buff: bytes.NewBuffer([]byte{}),
		logChan:make(chan []byte,5000),
		closeChan:make(chan interface{}),
		maxSize:1024,
	}
	go ws.run()
	return ws
}