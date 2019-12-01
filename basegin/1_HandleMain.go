package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.Default()也使用gin.New()来创建engine实例,但是会默认使用Logger和Recovery中间件
	// Logger中间件负责打印并输出日志,方便调试
	// Recovery中间件是如果程序执行过程中遇到panic中断了服务,则Recovery会恢复程序执行,并返回服务器500错误
	engine := gin.Default()
	//engine := gin.New()
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		// 获取接口请求地址
		path := context.FullPath()
		fmt.Println(path)
		// 获取get请求的参数,如果取不到使用默认值
		name := context.DefaultQuery("name", "默认值")
		fmt.Println(name)
		// 输出
		context.Writer.Write([]byte("Hello " + name))
	})
	engine.Handle("POST", "/login", func(context *gin.Context) {
		// 获取接口请求地址
		path := context.FullPath()
		fmt.Println(path)
		// 获取post请求的数据
		name := context.PostForm("name")
		password := context.PostForm("password")
		fmt.Println(name, password)
		// 输出
		context.Writer.Write([]byte("Login " + name))
	})
	// 指定端口
	engine.Run(":8080")

}
