package main

import (
	_ "gin-web/api"
	"gin-web/config"
	"gin-web/init/middleware"
	"gin-web/init/mysql"
	"gin-web/init/server"
	"gin-web/internal/router"
	"gin-web/pkg/filter"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// 加载中间件
	engine := middleware.Load(r, filter.Recover, filter.NoCache, filter.Options, filter.Secure, filter.RequestId(), filter.GinLogger())
	// 初始化路由
	router.InitRouter(engine)
	// 初始化 mysql 客户端
	mysql.InitMysql(config.AppConfig.Mysql.Username, config.AppConfig.Mysql.Password, config.AppConfig.Mysql.Host, config.AppConfig.Mysql.Port, config.AppConfig.Mysql.DbName)
	// 启动 http 服务器，
	server.Run(r, config.AppConfig.SC.Name, config.AppConfig.SC.Addr, nil)
}
