package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFile := "./test_write_fil.txt"
	writeFile := "./test_write.txt"

	data, err := ioutil.ReadFile(readFile)
	if err != nil {
		fmt.Println("读取文件失败", err)
	}
	err = ioutil.WriteFile(writeFile, data, 0666)
	if err != nil {
		fmt.Println("write file error:", err)
	}
}
