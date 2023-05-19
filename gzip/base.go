package gzip

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
)

func Decompress(d io.Reader, data interface{}) error {
	// 读取gzip压缩数据 /*Read gzip compressed data*/
	reader, err := gzip.NewReader(d)
	if err != nil {
		return err
	}
	defer reader.Close()
	// 将解压后的数据读取到字节数组中 /* Read the decompressed data into a byte array */
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	// fmt.Println(string(body))
	// 将字节数组数据映射到目标结构体中 /*Map byte array data into target structure*/
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

func Compression(jsonStr string) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	zw := gzip.NewWriter(&buffer)
	if _, err := zw.Write([]byte(jsonStr)); err != nil {
		return buffer, err
	}
	if err := zw.Close(); err != nil {
		return buffer, err
	}
	return buffer, nil
}
