package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//type Usb interface {
//	// 声明两个没有实现的方法
//	Start()
//	Stop()
//}
//
//type Phone struct {
//}
//
//func (p Phone) Start() {
//	fmt.Println("手机开始工作...")
//}
//
//func (p Phone) Stop() {
//	fmt.Println("手机停止工作...")
//}
//
//type Camera struct {
//}
//
//func (c Camera) Start() {
//	fmt.Println("相机开始工作...")
//}
//
//func (c Camera) Stop() {
//	fmt.Println("相机停止工作...")
//}
//
//type Computer struct {
//
//}
//
//// 接收一个Usb接口类型变量,只要实现了Usb接口 (所谓实现了Usb接口,就是指实现了Usb接口声明的所有方法)
//// 多态在python中通过方法来实现,在Go中通过接口来实现
//func (c Computer) Working(usb Usb) {
//	// 通过usb接口变量来调用start和stop方法
//	fmt.Println(reflect.TypeOf(usb))
//	usb.Start()
//	usb.Stop()
//}
//
//type AInterface interface {
//	Say()
//}
//
//type BInterface interface {
//	Hello()
//}
//// #############################################
//type Monster struct {
//
//}
//
//func (m Monster) Say() {
//	fmt.Println("monster say() ...")
//}
//
//func (m Monster) Hello() {
//	fmt.Println("monster hello() ...")
//}
//// #############################################
//
//// *****************************接口继承
//type CInterface interface {
//	AInterface
//	BInterface
//	test()
//}
//
//func (m Monster) test() {
//	fmt.Println("monster test() ...")
//}
//// *****************************
//func main() {
//	computer := Computer{}
//
//	phone := Phone{}
//	computer.Working(phone)
//
//	camera := Camera{}
//	computer.Working(camera)
//
//	// #############################################
//	// Monster实现了AInterface和BInterface
//	var monster Monster
//	var a1 AInterface = monster
//	var b1 BInterface = monster
//	a1.Say()
//	b1.Hello()
//	// #############################################
//
//	// *****************************
//	// Monster实现了CInterface中的所有方法
//	var c1 CInterface = monster
//	c1.Hello()
//	c1.Say()
//	c1.test()
//	// *****************************
//}