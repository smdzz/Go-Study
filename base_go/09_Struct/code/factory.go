package main

import (
	"fmt"
	"unsafe"
)

// 工厂模式

//不强制使用构造函数,首字母大写,其他包可以直接定义,不用通过工厂函数构建
type File struct {
	fd int // 文件描述符
	name string //文件名
}

// NewFile是File结构体类型对应的工厂方法,它返回一个指向结构体实例的指针
func NewFile(fd int, name string) *File {
	if fd > 0 {
		return nil
	}

	return &File{fd, name}
}


//强制使用构造函数,首字母小写,实例化该结构体的方式只能通过NewFile1实现
type file struct {
	fd int // 文件描述符
	name string //文件名
}

// NewFile1是file结构体类型对应的工厂方法,它返回一个指向结构体实例的指针
func NewFile1(fd int, name string) *file {
	if fd > 0 {
		return nil
	}

	return &file{fd, name}
}

func main() {
	// 查看结构体类型实例占用了多少内存
	size := unsafe.Sizeof(int8(3))
	fmt.Println(size)
	size1 := unsafe.Sizeof(NewFile1(100, "name"))
	fmt.Println(size1)
}