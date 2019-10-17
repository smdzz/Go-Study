//Go 语言切片是对数组的抽象。
//Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),
// 与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
//package main
//
//import "fmt"
//
//func main() {
//	var a = [3]int{2, 4, 6}
//	// b为切片
//	b := a[1:]
//	// append后面参数必须为切片,不能为数组,cap(b)为切片的容量,当len()=4,cap()=4时,再追加元素cap()=cap()*2,等cap()再满了之后还是乘以2
//	b = append(b, 42)
//	fmt.Println(len(b), cap(b))
//	b = append(b, 456)
//	fmt.Println(len(b), cap(b))
//	b = append(b, 446)
//	fmt.Println(len(b), cap(b))
//	fmt.Println(a, b)
//
//	// 切片初始化
//	var sli = [] int {1,2,3}
//	fmt.Println("sli: ", sli, len(sli), cap(sli))
//	// 容量cap() 10 为可选参数
//	s :=make([]int, 4, 10)
//	fmt.Println("s: ", s, len(s), cap(s))
//}

//package main
//
//import "fmt"
//
//func main() {
//	var numbers []int
//	// 空切片
//	printSlice(numbers)
//
//	if numbers==nil {
//		fmt.Printf("切片是空的")
//	}
//}
//
//func printSlice(x []int){
//	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}

//append() 和 copy() 函数
//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
//
//下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
//
//实例
package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	fmt.Println("hhhhhhhhhh")
	fmt.Println(numbers1, numbers)
	numbers2 := numbers[:]
	numbers[2] = 66666
	fmt.Println(numbers1, numbers, numbers2)
	printSlice(numbers1)
	a := [] int {1}
	printSlice(a)
	a = append(a, 2)
	printSlice(a)
	a = append(a, 3)
	printSlice(a)
	a = append(a, 4)
	printSlice(a)
	a = append(a, 5, 6, 7)
	printSlice(a)
}

func printSlice(x []int){
fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}