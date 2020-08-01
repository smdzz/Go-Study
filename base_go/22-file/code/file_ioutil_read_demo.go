package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// ioutil.ReadFile 一次性将问价读取到内存中,适用于文件不太大的情况
	content, err := ioutil.ReadFile("./test_file.txt")
	if err != nil {
		fmt.Println("read file error:", err)
	}
	// 输出读取到的内容
	fmt.Println(string(content))
}