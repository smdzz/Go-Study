package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

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