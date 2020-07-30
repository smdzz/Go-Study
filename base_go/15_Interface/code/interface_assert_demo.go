package main

import (
	"fmt"
)

type Usb interface {
	// 声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	PhoneNum string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

func (p Phone) Call() {
	fmt.Printf("给 %s 打电话...\n", p.PhoneNum)
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

func (c Camera) TakePhone() {
	fmt.Printf("使用 %s 拍照\n", c.Name)
}

type Computer struct {

}

func (c Computer) Working(usb Usb) {
	usb.Start()
	// 这里使用了类型断言,通过不同的类型去执行不同的方法
	switch usb.(type) {
	case Phone:
		usb.(Phone).Call()
	case Camera:
		usb.(Camera).TakePhone()
	default:
		fmt.Println("暂时不可识别的类型")
	}
	usb.Stop()
}

func main() {
	var computer Computer
	// 定义一个Usb接口数组,可以存放Phone和Camera结构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"110"}
	usbArr[1] = Phone{"120"}
	usbArr[2] = Camera{"詹姆的相机"}
	for _, usbIns := range usbArr {
		computer.Working(usbIns)
	}
}
