package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var register Register
		// 使用ShouldBind可以实现Post方式的提交数据的绑定工作
		if err := context.ShouldBind(&register);err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		context.Writer.Write([]byte(register.UserName + " Register"))
	})

	engine.Run()
}

type Register struct {
	UserName string `form:"name""`
	Phone 	 string `form:"phone""`
	PassWord string `form:"password""`
}
