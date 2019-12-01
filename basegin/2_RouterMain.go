package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//127.0.0.1:8080/hello?name=xiaoming
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		// Query为查询name参数,找不到为nil
		name := context.Query("name")
		context.Writer.Write([]byte("hello " + name))
	})

	engine.POST("login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		// GetPostForm结果为两个值,一个为取到得值,一个为是否取到值(布尔值)
		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println(username)
		}
		password, exists := context.GetPostForm("password")
		if exists {
			fmt.Println(password)
			context.Writer.Write([]byte("欢迎回来: " + username))
		}
	})

	engine.DELETE("/user/:id", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		// 获取 :id的值使用context.Param()
		id := context.Param("id")
		context.Writer.Write([]byte("删除用户的id为: " + id))
	})
	engine.Run()
}