package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("./html/*")
	// 前面参数为http请求去那个路径下寻找,后面参数表示本地工程的路径
	engine.Static("./img", "./img")
	engine.GET("/hellohtml", func(context *gin.Context) {
		fullpath := "请求路径为" + context.FullPath()
		context.HTML(http.StatusOK, "index.html", gin.H{"fullpath": fullpath, "title": "html测试"})
	})

	engine.Run()
}
