package main

import (
	"gin-web/config"
	"gin-web/init/mysql"
	"gin-web/internal/model/tabel"
	"gin-web/pkg/logs"
)

// 自动建表迁移
func main() {
	mysql.InitMysql(config.AppConfig.Mysql.Username, config.AppConfig.Mysql.Password, config.AppConfig.Mysql.Host, config.AppConfig.Mysql.Port, config.AppConfig.Mysql.DbName)
	err := mysql.DB.Self.AutoMigrate(&tabel.User{})
	if err != nil {
		logs.LG.Error(err.Error())
	}
	logs.LG.Info("自动迁移表成功")
}
