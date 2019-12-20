package main

import "fmt"

// go中有个特殊的概念就是标签,可以给for switch 或 select等流程控制代码块打上标签,配合标签标识符可以方便的跳转到某一个地方继续执行,有助于提高编程效率
// 首先来看break
func main() {
	//testBreak()
	//testBreakLabel()
	//testContinueLabel()
	testGoToLabel()
}

func testBreak() {
	for {
		x := 1
		switch {
		case x > 0:
			fmt.Println("A")
			//time.Sleep(1*time.Second)
			// 该处的break只是跳出了switch循环,并没有跳出for循环,所以这个程序会被一直执行下去,实验的时候使用time.sleep()叫他睡一会以免死机,手动哭脸
			break
		case x == 1:
			fmt.Println("B")
		default:
			fmt.Println("C")
		}
	}
}

func testBreakLabel() {
LOOP1:
	for i := 0; i < 5; i++{
		switch {
		case i == 1:
			fmt.Println("B")
		case i > 0:
			fmt.Println("A")
			// time.Sleep(1*time.Second)
			// 该处的break会跳出LOOP1的代码块之外,终止该程序
			continue LOOP1
		default:
			fmt.Println("C")
		}
	}
}
// print C B A

func testContinueLabel() {
LOOP1:
	for i := 0; i < 5; i++{
		switch {
		case i == 1:
			fmt.Println("B")
		case i > 0:
			fmt.Println("A")
			//time.Sleep(1*time.Second)
			// 该处的continue会跳到LOOP1的代码块位置继续执行任务,continue只能用于for循环
			continue LOOP1
		default:
			fmt.Println("C")
		}
	}
}
// print C B A A A

// goto只能在同一个函数中跳转,我们的无限for循环中当遇到i>2时,直接跳转打印break
func testGoToLabel() {
	var i int
	for {
		fmt.Println(i)
		i++
		if i > 2 {
			goto BREAK1
		}
	}
	BREAK1:
		fmt.Println("break")
}