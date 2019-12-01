package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.POST("/addstudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var person Person
		//gin框架提供了数据结构体和表单提交数据绑定的功能，提高表单数据获取的效率
		//
		if err := context.BindJSON(&person);err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Name)
		fmt.Println(person.Sex)
		fmt.Println(person.Age)
		context.Writer.Write([]byte("添加学生姓名: " + person.Name))
	})
	engine.Run()
}

type Person struct {
	// 通过tag标签的方式设置每个字段对应的form表单中的属性名，通过binding设置属性是否是必须。
	Name string `form:"name" binding:"required"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}