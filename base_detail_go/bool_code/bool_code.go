package main

import (
	"fmt"
	"os"
)

// 对于布尔值得好的命名能够提升代码的可读性,例如以is或者Is开头的isSorted,isFinished等,
// 使用这样的命名能够在阅读代码时获得与阅读正常语句一样的良好体验,标准库中的命名(unicode.IsDigit(ch)等)也是遵循这个原则的

var USER = os.Getenv("USER")
func main(){
	//var a bool // 直接声明bool类型,默认为false
	//a = 1 // 错误,应该为b=flase
	//a = bool(1) // 错误.应该为b = bool(false)
	var b bool
	b = (1 != 0)  // 编译时推断变量的值
	u := ("user" == USER)  // 运行时根据环境判断变量的值
	fmt.Println(b, u)  // 类型相同并且值相同才会返回false
}
