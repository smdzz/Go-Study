package main

import "fmt"

// 多重继承

// go语言通过在类型中嵌入所有必要的父类型,可以很简单的实现多重继承.
// go语言的多重集成不支持多重嵌套(即父类型内部不允许有匿名结构体字段)

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "拍照"
}

type Phone struct{}

func (c *Phone) Call() string {
	return "响铃"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("继承自Camera", cp.TakeAPicture())
	fmt.Println("继承自Phone", cp.Call())
}