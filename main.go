package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin_api/router"
)

func main() {
	// 解析配置文件
	// todo
	// 初始化 数据库
	// 设置模式
	// 创建gin引擎
	r := gin.Default()
	// bind router
	router.BindRouter(r)
	addr := ":8080"
	err := r.Run(addr)
	if err != nil {
		log.Fatal("run error, msg:", err.Error())
	}
}
