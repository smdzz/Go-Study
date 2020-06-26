//package main
//
//import "fmt"
//
//type Shaper interface {
//	Area() float32
//}
//
//type Square struct {
//	side float32
//}
//
//func (sq *Square) Area() float32 {
//	return sq.side * sq.side
//}
//
//func main() {
//	sq1 := new(Square)
//	sq1.side = 5
//	areaIntf := Shaper(sq1)
//	fmt.Printf("面积为: %f\n %T", areaIntf.Area(), areaIntf)
//}

//package main
//
//import "fmt"
//
//type Shaper interface {
//	Area() float32
//}
//
//type Square struct {
//	side float32
//}
//
//func (sq *Square) Area() float32 {
//	return sq.side * sq.side
//}
//
//type Rectangle struct {
//	length, width float32
//}
//
//func (r Rectangle) Area() float32 {
//	return r.length * r.width
//}
//
//func main() {
//	r := Rectangle{5, 8}
//	q := &Square{7}
//	shapers := []Shaper{r, q}
//	for n, _ := range shapers {
//      // 在调用shapers[n].Area()时, 只知道shapes[n]是一个Shaper对象,最后摇身一变成为了Square或Rectangle对象,并且表现出了相应的行为
//		fmt.Println("形状参数: ", shapers[n])
//		fmt.Println("形状面积是: ", shapers[n].Area())
//	}
//}

//package main
//
//import "fmt"
//
//type stockPosition struct {
//	ticker string
//	sharePrice float32
//	count float32
//}
//
//// 获取stock的值(价格)
//func (s stockPosition) getValue() float32 {
//	return s.sharePrice * s.count
//}
//
//type car struct {
//	make string
//	model string
//	price float32
//}
//
//// 获取car的值(价格)
//func (c car) getValue() float32 {
//	return c.price
//}
//
//// 定义具有价值的不同事物的'合同'
//type valuable interface {
//	getValue() float32
//}
//
//// 定义了一个使用valuable类型作为参数的函数showValue,所有实现了valuable的类型都可以使用这个函数
//func showValue(asset valuable) {
//	fmt.Printf("资产的价值是: %f\n", asset.getValue())
//}
//
//func main() {
//	var o valuable = stockPosition{"google", 855.3, 6}
//	showValue(o)
//	o = car{"AODI", "Q8", 880000}
//	showValue(o)
//}


package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Circle struct {
	radius float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	//areaIntf的类型是否是Square
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("areaIntf的类型是: %T\n", t)
	}

	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("areaIntf的类型是: %T\n", u)
	} else {
		fmt.Println("areaIntf不包含类型为Circle的变量")
	}

	fmt.Printf("面积为: %f\n %T", areaIntf.Area(), areaIntf)
}