package main

import "fmt"

// 一:
// 类型转换 type_name(expression)
// 注意事项: 要注意数据长度被截取而发生的数据精度损失(比如将浮点数强制转换为整数), 或值溢出问题(值超过转换的目标类型的值范围时)
//func main() {
//	sum := 23
//	count := 5
//	mean := float64(sum) / float64(count)
//	mean2 := sum / count
//	fmt.Printf("mean的值为:%f\n", mean)
//	fmt.Printf("mean的值为:%d\n", mean2)
//}


//二:
// 类型别名: byte是int8类型的别名,rune类型是int32类型的别名,go允许开发者自定义类型别名

//type (
//	字符串 string
//)
//
//func main() {
//	var b 字符串
//	b = "我是字符串"
//	fmt.Println(b)
//	a := "我也是字符串"
//	//fmt.Println(b + a) // 这里在编译器中是会报错的,因为类型别名与原类型是两种类型,不能直接进行操作,需要进行类型转换
//	fmt.Println(string(b)+a)
//}


// 三:
// 指针: 指针也是一个变量,它的值是另外一个变量的地址,也是存储器位置的直接地址,指针和变量一样,必须先声明一个指针,然后才能使用它来存储任意的变量地址
// 声明形式 var name *type

func main() {
	var ap *int
	c := 20
	ap = &c
	fmt.Printf("a的地址:%x\n", &c)
	fmt.Printf("ap的地址:%x\n", ap)
	fmt.Printf("*ao的值:%d\n", *ap)
	// nil指针
	var ptr *int
	fmt.Printf("ptr的值是:%x\n", ptr)  // 0
	a := 200
	b := 100
	fmt.Printf("交换之前a的值为%d\n", a)
	fmt.Printf("交换之前b的值为%d\n", b)
	swap(&a, &b)
	fmt.Printf("交换之后a的值为%d\n", a)
	fmt.Printf("交换之后b的值为%d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
