package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体,用于保存统计结果
type CharCount struct {
	EnCount    int // 记录英文的个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCOunt int // 记录其他字符的个数
}

func main() {
	// 统计一个文件中有多少个英文,多少个数字,多少个空格
	// 每读取一行,就去统计
	file, err := os.Open("test_type_count.txt")
	if err != nil {
		fmt.Println("open file err: ", err.Error())
	}

	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			count.analysis(str)
			break
		}
		count.analysis(str)
	}
	fmt.Printf("en:%v, num:%v, space:%v, other:%v\n", count.EnCount, count.NumCount, count.SpaceCount, count.OtherCOunt)
}

func (cc *CharCount)analysis(str string) {
	// 未适配中文,当有中文时放到了other中
	for _, v := range str {
		switch{
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			cc.EnCount++
		case v == ' ' || v == '\t':
			cc.SpaceCount++
		case v <= '9' && v >= '0':
			cc.NumCount++
		default:
			cc.OtherCOunt++
		}
	}
}
