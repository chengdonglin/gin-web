package mysql

import (
	"fmt"
	"gin-web/pkg/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func InitMysql(username string, password string, host string, port int, dbName string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.LG.Error(err.Error())
	}
	logs.LG.Info("Mysql 客户端初始化成功")
	conn := &Database{Self: db}
	DB = conn
}
