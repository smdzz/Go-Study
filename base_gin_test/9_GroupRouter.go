package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//// Logging to a file.日志不再打印到终端,全部存储到gin.log中
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	routerGroup := engine.Group("/user")
	routerGroup.POST("/register", register)

	routerGroup.POST("/login", login)

	routerGroup.DELETE("/:id", deleteHandle)

	engine.Run()
}


func register(context *gin.Context) {
	fullPath := "用户注册接口 " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func login(context *gin.Context) {
	fullPath := "用户登录接口 " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func deleteHandle(context *gin.Context) {
	fullPath := "用户删除接口 " + context.FullPath()
	fmt.Println(fullPath)
	id := context.Param("id")
	context.Writer.WriteString(fullPath + " " + id)
}