package main

import "github.com/gin-gonic/gin"

func main() {
	//这个是使用中间件的情况
	r := gin.New()

	//添加全局中间件到路由器 所有的请求都会被处理包括 404 405 静态资源等
	r.Use(gin.Logger())

	//拦截  localhost:8080/ping的get请求 处理方式是后面的func内容
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
