package file

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

const (
	BasePath = "./public/data/outputs/"
)

// 读取 CSV 文件
func ReadCsvFile(csvPath string, first bool) ([][]string, error) {
	var rows [][]string
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return rows, err
	}
	csvReader := csv.NewReader(csvFile)
	if !first {
		_, err := csvReader.Read()
		if err != nil {
			fmt.Println(err)
			return rows, err
		}
	}
	rows, err = csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return rows, err
	}
	return rows, nil
}

var UTF8BOM = []byte{239, 187, 191}

func hasBOM(in []byte) bool {
	return bytes.HasPrefix(in, UTF8BOM)
}

func stripBOM(in []byte) []byte {
	return bytes.TrimPrefix(in, UTF8BOM)
}

// 输出 csv
func WriteCSV(filename string, header []string, data [][]string) error {
	//递归创建目录··
	filesPaths, _ := filepath.Split(filename)
	os.MkdirAll(filesPaths, 0777)
	//递归创建目录··
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	file.WriteString("\xEF\xBB\xBF")
	if err := writer.Write(header); err != nil {
		fmt.Println(err)
	}
	for _, item := range data {
		if err := writer.Write(item); err != nil {
			fmt.Println(err)
		}
	}
	return err
}

func GetPathByfileName(filename string) string {
	var path = filename
	return path
}
