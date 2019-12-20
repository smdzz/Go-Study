package main

import "fmt"

// defer两大特点:
// 1. 只有当defer语句全部执行,defer所在的函数才算真正结束执行
// 2. 当函数中只有defer语句时,需要等待所有的defer语句执行完毕,才会执行return语句
// 作用:
// 由于defer的延迟特点,可以把defer语句用于回收资源 清理收尾等工作,例如打开数据库连接后或打开文件后的关闭

// defer 的本质是一个栈,遵循先入后出,后入先出的原则
//var num = 0
//func printNum() {
//	fmt.Println(num)
//}
//
//// 最后结果为5个5,为什么呢? 因为每次执行的时候将printNum()这个函数压入栈,此时函数内部还没有执行,当栈内函数弹出的时候,num=5,所以调用函数的时候返回值都是5
//func main() {
//	for ; num < 5; num++ {
//		defer printNum()
//	}
//}

var num = 0

func printNum(i int) {
	fmt.Println(i)
}
// 返回的结果是 4 3 2 1 0,这又是为什么呢?因为每次执行的时候将printNum(0)....printNum(4),依次压入栈中,
// 这时再弹出的时候执行每一个函数num值都是在压栈的同时就定好的,所以是倒序排列的
func main() {
	for ; num < 5; num++ {
		defer printNum(num)
	}
}