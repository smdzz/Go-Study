package main

import "fmt"

// for循环后面的三个字语句被我们称为:初始化子语句, 条件子语句, 后置子语句,这三者不能颠倒顺序,其中条件子语句是必须的
func main() {
	a := 5
	b := 0

	//for a > b {
	//	b++
	//	fmt.Printf("b的值是%d\n", b)
	//}

	// Go语言编译的时候会自动判断三个字语句中是否存在条件子语句
	for ; a > b ; {
		b++
		fmt.Printf("b的值是%d\n", b)
	}

	// 写成这样就会报错
	//for ; ; a > b{
	//	b++
	//	fmt.Printf("b的值是%d\n", b)
	//}
}
