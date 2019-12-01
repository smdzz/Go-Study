package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T){
	//第一种变量初始化方式
	//var a int = 1
	//var b int = 1
	//第二种变量初始化方式,在一个赋值语句中给多个变量赋值
	//var (
	//	a int = 1
	//	b int = 1
	//)
	//第三种变量初始化方式
	a:=1
	b:=1
	//fmt.Print(a)
	// 单元测试中可以使用t.Log(a)来打印变量
	t.Log(a)
	for i:=0;i<5;i++{
		//fmt.Print(" ", b)
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

// 交换两个变量的值
func TestExchange(t *testing.T){
	a := 1
	b := 2
	// 常见写法
	// tmp := a
	// a = b
	// b = tmp
	// 一个语句对多个变量赋值
	a, b = b, a
	t.Log(a, b)
}