//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}
////使用go开启线程
//func main() {
//	go say("我是go开启的线程")
//	say("我是主函数内被调用的函数")
//}

//通道,channel
package main

import (
	"fmt"
)

func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//time.Sleep(10000 * time.Millisecond)
	c <- sum // 把 sum 发送到通道 c
	fmt.Println("sum1 after channel pro")
}

func sum2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//time.Sleep(10000 * time.Millisecond)
	c <- sum // 把 sum 发送到通道 c
	fmt.Println("sum2 after channel pro")
}


func main() {
	//定义一个切片
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum1(s[:len(s)/2], c)
	go sum2(s[len(s)/2:], c)
	fmt.Println(1)
	x := <-c
	// 阻塞, 谁先到通道谁先赋值
	fmt.Println(1, x)
	y := <-c
	fmt.Println(1, y)
	//x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)
}


//// 通道的缓冲区
//package main
//
//import "fmt"
//
//func main() {
//	// 这里我们定义了一个可以存储整数类型的带缓冲通道
//	// 缓冲区大小为2
//	ch := make(chan int, 2)
//
//	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
//	// 而不用立刻需要去同步读取数据
//	ch <- 1
//	ch <- 2
//
//	// 获取这两个数据
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//}


//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//		time.Sleep(1000 * time.Millisecond)
//	}
//	//关闭通道并不会丢失里面的数据，只是让读取通道数据的时候不会读完之后一直阻塞等待新数据写入
//	close(c)
//	time.Sleep(2000 * time.Millisecond)
//	// 如果通道关闭,该函数还在继续执行,那么主函数运行完就会结束,
//	fmt.Println("dsajflk")
//}
//
//func main() {
//	c := make(chan int, 10)
//	go fibonacci(cap(c), c)
//	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
//	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
//	// 之后就结束了。如果上面的 c 通道不关闭，然后也没有别的操作,,就会形成死锁,上下都没有线程在运行,如果还有别的操作, 那么 range 函数就不
//	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
//
//	// 不加时间限制,说明上面是在通道放一个,下面的for取一个,当加上时间限制之后,上面函数会先将十个数加到通道里,然后直接一下全部取出
//	//time.Sleep(10000 * time.Millisecond)
//	for i := range c {
//		fmt.Println(i)
//	}
//}