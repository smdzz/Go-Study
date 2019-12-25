package main

import (
	"fmt"
)

func main() {
	fmt.Printf("年龄最大的是:%d\n", ageMinOrMax("max", 2, 4, 35, 65, 34))
	fmt.Printf("年龄最小的是:%d\n", ageMinOrMax("min", 2, 4, 35, 65, 34))
	f1([]int{2, 5, 9, 45, 34}...)
}

// 可变参数的使用(可变参数的原理是 可变参数的类型是一个切片)
func ageMinOrMax(m string, a ...int) int {
	if len(a) == 0 {
		return 0
	}
	switch m {
	case "min":
		min := a[0]
		for _, v := range a {
			if v < min {
				min = v
			}
		}
		return min
	case "max":
		max := a[0]
		for _, v := range a {
			if v > max {
				max = v
			}
		}
		return max
	default:
		return -1
	}
}
// 变长参数可以作为对应类型的slice进行二次传递,从内部实现机制上来说,类型...type本质上是一个数组切片,也就是[]type
// []int{2,3,4}...   后面的...会把数组打散拆成一个一个的数字
func f1(arr ...int){
	f2(arr...)
	fmt.Println("")
	f3(arr)
}

func f2(arr ...int) {
	for _, char := range arr {
		fmt.Printf("%d\n", char)
	}
}

func f3(arr []int) {
	for _, char := range arr {
		fmt.Printf("%d\n", char)
	}
}