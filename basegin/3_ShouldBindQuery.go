package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var student Student
		// 使用ShouldBindQuery可以实现Get方式的数据请求的绑定
		// ShouldBindQuery()返回的是一个err类型,将context中的数据绑定到Student这个类型,如果绑定成功,err=nil,如果失败,nil有值
		err := context.ShouldBindQuery(&student)
		if err != nil {
			fmt.Println(err)
			fmt.Println("cuole")
			log.Fatal(err.Error())
			return
		}
		fmt.Println(student.Name)
		fmt.Println(student.Classes)
		context.Writer.Write([]byte("hello: " + student.Name))
	})

	engine.Run()
}

type Student struct {
	Name 	string `form:"name"`
	Classes string `form:"classes"`
}