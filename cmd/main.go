package main

import (
	_ "gin-web/api"
	"gin-web/config"
	"gin-web/init/server"
	"gin-web/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	server.Run(r, config.AppConfig.SC.Name, config.AppConfig.SC.Addr, nil)
}
