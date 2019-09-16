package main

import "github.com/gin-gonic/gin"

func main() {
	//这个是默认情况 不是用中间件 middleWare的情况下
	r := gin.Default()
	//拦截  localhost:8080/ping的get请求 处理方式是后面的func内容
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
