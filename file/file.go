package file

import (
	"io/ioutil"
	"os"
)

// 创建目录
func CreateFileDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//获取目录下所有文件名称
func GetFiles(folder string) []string {
	var fileArr []string
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			GetFiles(folder + "/" + file.Name())
		} else {
			fileArr = append(fileArr, file.Name())
		}
	}
	return fileArr
}
