package main

import (
	"go-gin/config"
	"go-gin/middleware"
	"go-gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	//var port string
	//flag.StringVar(&port, "port", "", "port")
	//flag.Parse()
	//log.Printf("port:%s", port)
	//SetUp(port)
	SetUp()
}

func SetUp() {
	server := gin.Default()

	// 读取配置文件
	config.InitConfig()

	//mysql.InitMysql()

	// 日志
	//server.Use(middleware.InterfaceLogger())
	// 跨域
	server.Use(middleware.CorsMiddleware())

	// 初始化router
	router.InitRouter(server)

	// 启动
	//if port == "" {
	//	port = config.GetServerConfig().Port
	//}
	if err := server.Run(":" + config.GetServerConfig().Port); err != nil {
		panic(err)
	}
}

type Logger interface {
}
