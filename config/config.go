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
	GC    *GrpcConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name string
	Addr string
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
	conf.ReadGrpcConfig()
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

func (c *Config) ReadGrpcConfig() {
	gc := &GrpcConfig{
		Name: c.viper.GetString("grpc.name"),
		Addr: c.viper.GetString("grpc.addr"),
	}
	c.GC = gc
}
