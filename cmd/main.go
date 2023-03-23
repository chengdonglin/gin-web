package main

import (
	_ "gin-web/api"
	"gin-web/config"
	"gin-web/init/mysql"
	"gin-web/init/server"
	"gin-web/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化路由
	router.InitRouter(r)
	// 初始化 mysql 客户端
	mysql.InitMysql(config.AppConfig.Mysql.Username, config.AppConfig.Mysql.Password, config.AppConfig.Mysql.Host, config.AppConfig.Mysql.Port, config.AppConfig.Mysql.DbName)
	// 启动 http 服务器，
	server.Run(r, config.AppConfig.SC.Name, config.AppConfig.SC.Addr, nil)
}
