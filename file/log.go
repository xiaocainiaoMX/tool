package file

import (
	"github.com/btccom/admin_background/application/library/tool"
	"log"
	"os"
	"path/filepath"
	"time"
)

type fileLog struct {
	FilePath     string
	FileName     string
	FilePathName string
}

var FileLog fileLog

func init() {
	FileLog.FilePath = "./logs/"
	FileLog.FileName = time.Now().Local().Format(tool.UTCDay) + ".log"
	FileLog.FilePathName = FileLog.FilePath + FileLog.FileName
}

func (fl *fileLog) SetFilePathName(fileName string) *fileLog {
	fl.FilePathName = fileName
	//递归创建目录··
	filesPaths, _ := filepath.Split(fl.FilePathName)
	err := os.MkdirAll(filesPaths, 0777)
	if err != nil {
		panic(err)
	}
	return fl
}
func (fl *fileLog) SetFilePath(path string) *fileLog {
	fl.FilePathName = path
	//递归创建目录··
	filesPaths, _ := filepath.Split(fl.FilePathName)
	err := os.MkdirAll(filesPaths, 0777)
	if err != nil {
		panic(err)
	}
	fl.FilePathName = fl.FilePathName + fl.FileName
	return fl
}
func (fl *fileLog) SetFileName(name string) *fileLog {
	fl.FileName = name
	fl.FilePathName = fl.FilePathName + fl.FileName
	return fl
}

func (fl *fileLog) AddLog(msg string) *fileLog {
	//创建日志文件
	f, err := os.OpenFile(fl.FilePathName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	// 写入日志内容
	log.Println(msg)
	return fl
}
