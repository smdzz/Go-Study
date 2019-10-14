// 输出hello world!
// 表示hello.go文件所在的包是main,在go中,每个文件都必须归属于一个包
package main
// 表示引入一个包,包名为fmt,引入该包后,就可以使用fmt包的函数,比如,fmt.Println()
import (
	"fmt"
	"os"
)
// func是一个关键字,表示一个函数
// main是函数名,是一个主函数,即程序的入口
func main() {
	// 如果参数大于1,打印出第1位的参数
	if len(os.Args)>1 {
		fmt.Println("hello world!", os.Args[1])
	}
	// 打印命令行参数
	fmt.Println(os.Args)
	// 表示调用fmt.Println函数输出hello world!
	fmt.Println("hello world!")
	// 退出主函数,使用异常退出
	os.Exit(-1)
	// 如果参数大于1,打印参数
}
// 使用go build hello.go进行编译
// 使用go run hello.go直接执行(在底层也进行编译)
