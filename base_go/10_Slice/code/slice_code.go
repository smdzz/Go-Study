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

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50}
	// 创建索引时,还可以用来控制新切片的容量,其目的不是要新增容量,而是要限制容量
	//对于slice[i:j:k](2:3:4),长度为j-k(即3-2=1),容量为k-i(4-2=2)
	ss := slice[2:3:4]
	fmt.Println(ss)
	s := slice[2:3]
	fmt.Println(s)
	// 由于s切片是对原切片的引用,切片间共享同一个底层数组,当进行append操作并且不会大于切片的容量时,会修改原切片的数据,
	// 当进行append操作并且大于切片的容量时,会创建一个新的容量更大的切片,触发数据搬移操作,此时与原切片不再共享同一个底层数组,不会再影响原切片的值
	s = append(s, 99, 1000, 434)
	fmt.Println(s)
	fmt.Println(slice)
	// 将一个切片追加到另外一个切片中,使用...运算符,...会将切片中的数据拿出来
	s = append(s, []int{3, 5, 7}...)

	// 多维数组
	slices := [][]int{{100}, {10, 20}}
	slices[0] = append(slices[0], 33)
	fmt.Println(slices)

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