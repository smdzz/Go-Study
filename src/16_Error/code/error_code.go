//package main
//
//import (
//	"fmt"
//)
//
//// 定义两个数的结构体
//type DivideError struct {
//	dividee int
//	divider int
//}
//
//// 实现DivideError的Error接口方法
//func (de *DivideError) Error() string {
//	strFormat := `
//   Cannot proceed, the divider is zero.
//   dividee: %d
//   divider: 0
//`
//	return fmt.Sprintf(strFormat, de.dividee)
//}
//
//// 定义两个int类型进行除法计算的函数
//func Devide(vardividee int, vardivider int) (result int, errMsg string) {
//	if vardivider == 0 {
//		ddata := DivideError{
//			dividee: vardividee,
//			divider: vardivider,
//		}
//		errMsg = ddata.Error()
//		return
//	} else {
//		result = vardividee / vardivider
//		return result, ""
//	}
//}
//
//// 两数相加
//func addDevide(vardividee int, vardivider int) (result int) {
//	return vardivider + vardividee
//}
//
//func main() {
//	// 两数相除
//	if result, errMsg := Devide(100, 10); errMsg=="" {
//		fmt.Println("100 / 10 = ", result)
//	}
//	if result, errMsg := Devide(100, 0); errMsg!="" {
//		fmt.Println(result, "errMsg = ", errMsg)
//	}
//	fmt.Println(addDevide(100, 200))
//}


// 详解 panic defer recover异常处理
package main

import (
	"errors"
	"fmt"
)

func main() {
	testError()
	fmt.Println("当上面函数的panic被recover()捕获后,我就可以继续执行了")
}

func testError() {
	//defer 类似于python里面的finally,最后执行
	defer func() {
		fmt.Println("我是defer")
		// recover 类似于python中的except,用来捕获异常,保证程序继续执行,但容易形成僵尸进程(还在启动着,但是已经无法提供服务)
		if err := recover(); err != nil {
			fmt.Println("recover from ", err)
		}
	}()
	fmt.Println("start")
	// panic主动抛出异常
	panic(errors.New("我这里出现错误,运行不下去了"))
	//os.Exit(1)
}
