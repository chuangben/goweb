package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//1 加载配置文件

	//2 初始化日志

	//3 初始化mysql

	//4 初始化redis

	//5 注册路由

	//6 启动服务
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8081")
}
