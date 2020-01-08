//数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
package main

//import "fmt"
//
//var name = [...]string{"张三", "李四"}
//func main() {
//	var number = [5]int{23, 34, 45}
//	var a = [3][4]int{
//		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
//		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
//		{8, 9}}   /* 第三行索引为 2 */
//	fmt.Println(name)
//	fmt.Println(name[0])
//	fmt.Println(number[4])
//	fmt.Println(a)
//	fmt.Println(a[2])
//	fmt.Println(a[2][3])
//}

import "fmt"

func main() {
	/* 数组长度为 5 */
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f\n", avg)

	// 知道数组的长度给特定的下标指定具体值
	arrayString := [5]string{"墨尔本", 1: "中国", "电话费", 4: "日本"}
	fmt.Println(arrayString)

	// 声明一个指针数组
	a := 100
	arrayPoint := [5]*int{0: new(int), 3: new(int), 4: new(int)}
	arrayPoint[0] = &a
	fmt.Println(arrayPoint)

	diffArray()

	diffPointArray()

	multidimensionalArray()
}
func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}

// 数组赋值
func diffArray() {
	// 数组是一个类型值,数组也可以像函数一样用在赋值操作中,变量名代表整个数组,因此同样类型的数组可以复制给另一个数组
	// 数组作为一个变量类型,它包括数组长度和每个元素的类型两个部分,只有这两个类型都相同的数组才是类型相同的数组,才能相互赋值
	var array1 [5]string
	array2 := [5]string{"中国", "日本", "美国", "意大利", "新加坡"}
	array1 = array2
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array1 == array2)
	//相互赋值之后两个数组还是独立的,修改其中一个不会影响其中另外一个
	array2[0] = "英国"
	fmt.Println(array1, array2)
	fmt.Println(array1 == array2)
}

// 数组指针赋值
func diffPointArray() {
	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}
	*array2[0] = "中国"
	*array2[1] = "美国"
	*array2[2] = "俄罗斯"
	array1 = array2
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array1 == array2)
	//如果复制数组指针,只会复制指针的值,而不会复制指针指向的值,所以修改指针数组中某一个指针指向的值,两个数组中对应指针指向的值都会更改
	*array1[0] = "英国"
	fmt.Println(array1, array2)
	fmt.Println(array1 == array2)
}

//多维数组赋值
func multidimensionalArray() {
	// 声明两个不同的二位整形数组
	var arr1 [2][2]int
	var arr2 [2][2]int
	// 为每个元素赋值
	arr2[0][0] = 10
	arr2[0][1] = 20
	arr2[1][0] = 30
	arr2[1][1] = 40
	// 将arr2的值赋值给arr1
	arr1 = arr2

	fmt.Println(arr1, arr2)
}
