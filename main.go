package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kanogachi/gin/ginday01/dbconn"
	"github.com/kanogachi/gin/ginday01/routers"
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Panicf("%s,error:%v\n", msg, err)
	}
}

func main() {
	// 初始化数据库
	err := dbconn.InitDB()
	failOnError("Initial database failed", err)
	// 服务运行结束后关闭数据库连接（可选）
	defer dbconn.Database.Close()
	// 初始化路由
	router := gin.Default()
	// 注册路由
	routers.RegisterRouter(router)
	// 开启服务
	err = router.Run(":8080")
	failOnError("Start server failed", err)
}
