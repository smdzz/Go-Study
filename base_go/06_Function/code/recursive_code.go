package main

import "fmt"

// 递归函数

func main() {
	// new(Type)返回的是指针的地址
	d := new(int)
	fmt.Println(d)
	var i uint64 = 20
	fmt.Printf("%d的阶乘是%d", i, Factorial(i))
	for i := 0; i < 20; i++ {
		fmt.Println(Fibonacci(i))
	}
}

// 使用递归求阶乘
func Factorial(n uint64)(result uint64) {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// 使用递归求斐波那契数
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}
