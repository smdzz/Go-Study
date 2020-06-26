package main
// 继承
import (
	"fmt"
	"math"
)

// Point是一个已知的具名结构体
type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// NamePoint结构体内部包含匿名字段Point
type NamePoint struct {
	Point
	name string
}


func main() {
	n := NamePoint{Point{3, 4}, "Python"}
	// 将一个已存在类型的字段和方法注入到另一个类型里极为内嵌,匿名字段上的方法"晋升"成为了外层类型的方法
	fmt.Println(n.Abs())
}

// 具名结构体也可以有只作用于本身实例而不作用于内嵌"父"类型上的方法,即覆写方法,此时再调用n.Abs()返回5000
//func (n NamePoint) Abs() float64 {
//	return n.Point.Abs() * 1000
//}
