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
	var  balance = [5]int {1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage( balance, 5 ) ;

	/* 输出返回的平均值 */
	fmt.Printf( "平均值为: %f ", avg );
}
func getAverage(arr [5]int, size int) float32 {
	var i,sum int
	var avg float32

	for i = 0; i < size;i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
