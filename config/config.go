package config

import (
	"gin-web/pkg/logs"
	"github.com/spf13/viper"
	"log"
	"os"
)

var AppConfig = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	Mysql *MysqlConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{
		viper: v,
	}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yml")
	conf.viper.AddConfigPath(workDir + "/config") // 可以添加多个
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.ReadMysqlConfig()
	return conf
}

func (c *Config) ReadServerConfig() {
	SC := &ServerConfig{}
	SC.Name = c.viper.GetString("server.name")
	SC.Addr = c.viper.GetString("server.addr")
	c.SC = SC
}

func (c *Config) InitZapLog() {
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("zap.maxSize"),
		MaxAge:        c.viper.GetInt("zap.maxAge"),
		MaxBackups:    c.viper.GetInt("zap.maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadMysqlConfig() {
	m := &MysqlConfig{
		Username: c.viper.GetString("mysql.username"),
		Password: c.viper.GetString("mysql.password"),
		Host:     c.viper.GetString("mysql.host"),
		Port:     c.viper.GetInt("mysql.port"),
		DbName:   c.viper.GetString("mysql.dbname"),
	}
	c.Mysql = m
}
