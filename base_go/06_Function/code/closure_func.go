package main

import "fmt"

func main() {
	var f = adder()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

// 在多次调用中, 变量x的值是被保留的, 即0+1=1, 1+2=3, 3+3=6,闭包函数保存并积累其中的变量的值,不管外部函数是否退出,他都能继续操作外部函数中的局部变量
func adder() func(int) int {
	var x int
	return func (d int) int {
		fmt.Println("x = ", x, "d = ", d)
		x += d
		return x
	}
}

