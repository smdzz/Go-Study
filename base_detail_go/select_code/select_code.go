package main

import "fmt"

// select选择语句,用于配合通道(channel)的读写操作,用于多个channel的并发读写操作
func main() {
	a := make(chan int, 1024)
	b := make(chan int, 1024)

	for i := 0; i< 10; i++ {
		fmt.Printf("第%d次", i)
		a <- 1
		b <- 1
		select {
		// 同时在a和b中选择,哪个有内容就从那个读,由于channel中读写操作是阻塞操作,使用select语句可以避免单个channel的阻塞,
		// 此外select同样可以使用default代码块,用于避免所有channel同时阻塞
		case <-a:
			fmt.Println("from a")
		case <-b:
			fmt.Println("from b")
		}
	}
}
