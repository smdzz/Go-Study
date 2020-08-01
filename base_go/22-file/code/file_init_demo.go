package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	// 1. file叫file对象
	// 2. file叫file指针
	// 3. file叫file文件句柄
	file, err := os.Open("../note.txt")
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}
	fmt.Printf("file=%v", file)
	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err: ", err.Error())
	}

}
