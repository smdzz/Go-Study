package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// 调用JSON将map类型的数据转换成为json格式并返回给前端，第一个参数200表示设置请求返回的状态码。和http请求的状态码一致。
	engine.GET("/hellojson", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"message": "我是json返回数据",
			"data": "将map格式返回json",
		})
	})

	// 结构体也是可以直接转换为JSON格式进行返回
	engine.GET("/jsonstruct", func(context *gin.Context) {
		rep := Response{Code: 1,Message: "我是json返回数据",Data: "将结构体转换为json返回"}
		context.JSON(200, &rep)
	})
	engine.Run()
}

type Response struct {
	Code int
	Message string
	Data interface{}
}