package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	
	engine.GET("/returnbyte", func(context *gin.Context) {
		fullpath := "我是本次的请求地址: " + context.FullPath()
		fmt.Println(fullpath)
		context.Writer.Write([]byte(fullpath))
	})

	engine.GET("/returnstring", func(context *gin.Context) {
		fullpath := "我是本次的请求地址: " + context.FullPath()
		fmt.Println(fullpath)
		context.Writer.WriteString(fullpath)
	})
	
	engine.Run()
}
