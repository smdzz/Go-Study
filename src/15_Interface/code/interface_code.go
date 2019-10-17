package main

import "fmt"

type DogClass interface {
	changeName(dog Dog, name string)
	changeColor(color string)
}

type Dog struct {
	name string
	age int
	color string
}

// 想要修改Dog的属性,必须要传入dog的指针结构体,才可以进行修改
// (dog *Dog)声明属于哪个结构体, changeName为上面定义好的函数,  后面的string为返回值类型
func (dog *Dog) changeName(name string) string {
	dog.name = name
	return dog.name
}

func main() {
	dog1 := new(Dog)
	fmt.Println(dog1.changeName("xiogou"))
	fmt.Println(dog1)
}
