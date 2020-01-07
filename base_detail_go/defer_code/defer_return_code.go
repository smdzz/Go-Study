package main

import "fmt"

func main() {
	fmt.Println("return:", a())
	fmt.Println("return:", b())
	d()
}

// 带有defer的return的实现逻辑
// 1. 第一步是先给返回值赋值(若为有名函数直接赋值,若为匿名函数则先声明再赋值)
// 2. 第二步调用RET返回指令并传入返回值,而RET会检查defer是否存在,若存在就先逆序插播defer语句(RET返回指令就是Return)
// 3. 最后RET携带返回值退出函数
// 未定义返回值
func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

func c() {
	defer un(trace("c"))
	fmt.Println("c的逻辑代码")
}

func d() {
	defer un(trace("d"))
	fmt.Println("d的逻辑代码")
	c()
}

func trace(s string) string {
	fmt.Println("开始执行", s)
	return s
}

func un(s string) {
	fmt.Println("结束执行", s)
}

//开始执行 d
//d的逻辑代码
//开始执行 c
//c的逻辑代码
//结束执行 c
//结束执行 d