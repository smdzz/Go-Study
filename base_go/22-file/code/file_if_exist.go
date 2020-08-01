package main

import (
	"fmt"
	"os"
)

func main() {
	exist, err := PathExist("./test_write.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if exist {
		fmt.Println("该文件存在")
		return
	}
	fmt.Println("该文件不存在")
	return
}

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	// 如果报错为nil,则说明文件或者目录存在
	if err == nil {
		return true, nil
	}

	// 如果返回的错误类型使用os.IsNotExist判断为true,则说明文件或者文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}

	// 如果返回的错误为其他类型,则不确定是否存在
	return false, err
}
