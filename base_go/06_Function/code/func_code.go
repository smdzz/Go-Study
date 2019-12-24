package main


//// 函数调用
//func main() {
//	/* 定义局部变量 */
//	a := 100
//	b := 200
//	var ret int
//	/* 调用函数并返回最大值 */
//	ret = max(a, b)
//
//	fmt.Printf( "最大值是 : %d\n", ret )
//}
//
///* 函数返回两个数的最大值 */
//func max(num1, num2 int) int {
//	/* 定义局部变量 */
//	var result int
//	if num1 > num2 {
//		result = num1
//	} else {
//		result = num2
//	}
//	return result
//}


//// 函数返回多个值
//func swap(x, y string) (string, string) {
//	return y, x
//}
//
//func main() {
//	a, b := swap("Google", "Runoob")
//	fmt.Println(a, b)
//}


// 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
//默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

//import "fmt"
//
//func main() {
//	/* 定义局部变量 */
//	var a int = 100
//	var b int = 200
//
//	fmt.Printf("交换前 a 的值为 : %d\n", a )
//	fmt.Printf("交换前 b 的值为 : %d\n", b )
//
//	/* 通过调用函数来交换值 */
//	swap(a, b)
//
//	fmt.Printf("交换后 a 的值 : %d\n", a )
//	fmt.Printf("交换后 b 的值 : %d\n", b )
//}
//
///* 定义相互交换值的函数 */
//func swap(x, y int) int {
//	var temp int
//
//	temp = x /* 保存 x 的值 */
//	x = y    /* 将 y 值赋给 x */
//	y = temp /* 将 temp 值赋给 y*/
//
//	return temp;
//}


// 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
//引用传递指针参数传递到函数内

//import "fmt"
//
//func main() {
//	/* 定义局部变量 */
//	var a int = 100
//	var b int= 200
//
//	fmt.Printf("交换前，a 的值 : %d\n", a )
//	fmt.Printf("交换前，b 的值 : %d\n", b )
//
//	/* 调用 swap() 函数
//	 * &a 指向 a 指针，a 变量的地址
//	 * &b 指向 b 指针，b 变量的地址
//	 */
//	swap(&a, &b)
//
//	fmt.Printf("交换后，a 的值 : %d\n", a )
//	fmt.Printf("交换后，b 的值 : %d\n", b )
//}
//
//func swap(x *int, y *int) {
//	var temp int
//	temp = *x  /* 保存 x 地址上的值 */
//	*x = *y      /* 将 y 值赋给 x */
//	*y = temp    /* 将 temp 值赋给 y */
//}


//Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。以下实例中我们在定义的函数中初始化一个变量，该函数仅仅是为了使用内置函数 math.sqrt()，实例为：

//import (
//"fmt"
//"math"
//)
//
//func main(){
//	/* 声明函数变量 */
//	getSquareRoot := func(x float64) float64 {
//		return math.Sqrt(x)
//	}
//
//	/* 使用函数 */
//	fmt.Println(getSquareRoot(9))
//
//}


//import "fmt"
//
//// 声明一个函数类型
//type cb func(int) int
//
//func main() {
//	testCallBack(1, callBack)
//	testCallBack(2, func(x int) int {
//		fmt.Printf("我是回调，x：%d\n", x)
//		return x
//	})
//}
//
//func testCallBack(x int, f cb) {
//	f(x)
//}
//
//func callBack(x int) int {
//	fmt.Printf("我是回调，x：%d\n", x)
//	return x
//}

//Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
//import "fmt"
//func main() {
//	add_func := add(1,2)
//	fmt.Println(add_func(1,1))
//	fmt.Println(add_func(0,0))
//	fmt.Println(add_func(2,2))
//}
//// 闭包使用方法
//func add(x1, x2 int) func(x3 int,x4 int)(int,int,int)  {
//	i := 0
//	return func(x3 int,x4 int) (int,int,int){
//		i++
//		return i,x1+x2,x3+x4
//	}
//}


//Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：
//
//func (variable_name variable_data_type) function_name() [return_type]{
//	/* 函数体*/
//}
//下面定义一个结构体类型和该类型的一个方法：

import "fmt"

/* 定义结构体 */
type Circle struct {
	radius float64
}

//func main() {
//	var c1 Circle
//	c1.radius = 10.00
//	fmt.Println("圆的面积 = ", c1.getArea())
//}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func pipe(ff func() int) int {
	return ff()
}

type FormatFunc func(s string, x, y int) string

func format(ff FormatFunc, s string, x, y int) string {
	return ff(s, x, y)
}

func main() {
	s1 := pipe(func() int {return 100})
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	fmt.Println(s1, s2)
}