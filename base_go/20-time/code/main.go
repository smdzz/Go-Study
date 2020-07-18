package main

import (
	"fmt"
	"time"
)

func main(){
	// 获取当前时间
	nowTime := time.Now()
	fmt.Println(nowTime)
	fmt.Println(time.Now())
	fmt.Println(nowTime.Year())
	fmt.Println(nowTime.Month())
	fmt.Println(int(nowTime.Month()))
	fmt.Println(nowTime.Day())
	fmt.Println(nowTime.Hour())
	fmt.Println(nowTime.Minute())
	fmt.Println(nowTime.Second())
	fmt.Println(nowTime.Format("2006-01-02 15:04:05"))
	//time.Sleep(1 * time.Second)
	fmt.Println(nowTime.Unix())
	fmt.Println(nowTime.UnixNano())
	fmt.Println(time.Unix(1000, 100000))
	fmt.Println(nowTime.UTC())
	fmt.Println(int(time.Now().Sub(nowTime)))
}
