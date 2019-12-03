package main

import "fmt"

func main() {
	str := "abc中国"
	//返回的char的类型为rune(int32)类型,可以使用string(char)来取到字符串的值
	for i, char := range str {
		fmt.Printf("字符串第%d个字符的值为%v\n", i, char)
		//字符串第0个字符的值为97
		//字符串第1个字符的值为98
		//字符串第2个字符的值为99
		//字符串第3个字符的值为20013
		//字符串第6个字符的值为22269
	}

	// 忽略第一个值,只关心返回的单个字串
	for _, char := range str {
		fmt.Println(string(char))
	}

	// 忽略第二个值,只返回索引
	for i := range str {
		fmt.Println(i)
	}

	for range str {
		fmt.Println("忽略全部的值,只执行我这个代码块")
	}

	// map类型使用range遍历的第一个值为key, 第二个值为value;
	mapTest := map[string]int{"a": 1, "b": 2}
	for k, v := range mapTest {
		fmt.Println(k, v)
	}

	// array/slice使用range遍历第一个值为index,第二个值为索引对应的值
	numbers := []int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Printf("第%d次,x的值为%d. \n", i, x)
	}

	//
}