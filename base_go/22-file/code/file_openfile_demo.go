package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ReWrite()
	// flag为只写
	fs, err := os.OpenFile("./test_write_fil.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	defer fs.Close()
	// 准备写入的语句为"hello, gardon\n"
	str := "hello, gardon\n"
	// 写入时使用带缓存的*writer
	writer := bufio.NewWriter(fs)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存的,因此在调用writerString方法时,其实是先写入到缓存中,flush的时候才将缓存区的数据写入到文件
	writer.Flush()

	appendWrite()
}

func ReWrite() {
	// flag为清空文件并只写
	fs, err := os.OpenFile("./test_write_fil.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	defer fs.Close()
	// 准备写入的语句为"hello, gardon\n"
	str := "我将原来的文件覆盖了\n"
	// 写入时使用带缓存的*writer
	writer := bufio.NewWriter(fs)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存的,因此在调用writerString方法时,其实是先写入到缓存中,flush的时候才将缓存区的数据写入到文件
	writer.Flush()
}

func appendWrite() {
	// flag为向已有的文件追加
	fs, err := os.OpenFile("./test_write_fil.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	defer fs.Close()
	// 准备写入的语句为"hello, gardon\n"
	str := "我是新来的,只能放最后\n"
	// 写入时使用带缓存的*writer
	writer := bufio.NewWriter(fs)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存的,因此在调用writerString方法时,其实是先写入到缓存中,flush的时候才将缓存区的数据写入到文件
	writer.Flush()
}