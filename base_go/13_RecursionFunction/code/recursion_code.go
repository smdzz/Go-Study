package main

import "fmt"

//func main() {
//	fmt.Println(Factorial(uint64(15)))
//}
//
//
//func Factorial(num uint64)(result uint64) {
//
//	if num > 1 {
//		// 递归相乘
//		result = num * Factorial(num-1)
//		return  result
//	} else {
//		// 递归终止条件
//		return 1
//	}
//}

func main(){
	fmt.Println(fibonacci(4))
	// 求前20位斐波那契数列
	for i:=1; i<=20;i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}


func fibonacci(num int)(result int) {
	if num < 2 {
		return num
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}