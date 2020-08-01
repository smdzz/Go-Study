package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 这里拷贝图片是可以的
	//a, err := CopyFile("./1.png", "/home/mengd/图片/1.png")
	err := CopyFile("./1.txt", "./test_file.txt")
	if err != nil {
		fmt.Println("拷贝文件失败")
		return
	}
	fmt.Println("拷贝文件成功")
	return
}

func CopyFile(dstFileName string, srcFileName string) (err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file error: ", err.Error())
		return
	}
	defer srcFile.Close()
	// 通过srcFile,获取到Reader
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file error: ", err.Error())
		return
	}

	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	// 拷贝文件是空的,看io.Copy中没有flush,这里copy完需要手动flush,否则拷贝文件是没有内容的
	_, err = io.Copy(writer, reader)
	if err == nil {
		writer.Flush()
	}
	return err
}