package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test_file.txt")
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}

	// 关闭文件,一般使用defer保证文件句柄一定会被关闭
	defer file.Close()
	// 创建一个reader,是带缓冲的
	//const (
	//	defaultBufSize = 4096
	//)默认的缓冲区是4096个字节
	// 适合于一个比较大的文件的情况
	reader := bufio.NewReader(file)
	for {
		// 读到一个换行就结束
		str, err := reader.ReadString('\n')
		// io.EOF表示文件的末尾
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
}