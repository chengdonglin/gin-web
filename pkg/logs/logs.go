package logs

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var LG *zap.Logger

type LogConfig struct {
	DebugFileName string `json:"debugFileName"`
	InfoFileName  string `json:"infoFileName"`
	WarnFileName  string `json:"warnFileName"`
	MaxSize       int    `json:"maxSize"`
	MaxAge        int    `json:"maxAge"`
	MaxBackups    int    `json:"maxBackups"`
}

// InitLogger 初始化logger
func InitLogger(cfg *LogConfig) error {
	writeSyncerDebug := getLogWrite(cfg.DebugFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	writeSyncerInfo := getLogWrite(cfg.InfoFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	writeSyncerError := getLogWrite(cfg.WarnFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	// 文件输出
	debugCore := zapcore.NewCore(encoder, writeSyncerDebug, zapcore.DebugLevel)
	infoCore := zapcore.NewCore(encoder, writeSyncerInfo, zapcore.InfoLevel)
	warnCore := zapcore.NewCore(encoder, writeSyncerError, zapcore.ErrorLevel)
	// 标准输出
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	std := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	core := zapcore.NewTee(debugCore, infoCore, warnCore, std)
	LG = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(LG) //替换全局的logger实例，后续在其他地方只需要使用zap.L()调用即可
	return nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWrite(fileName string, maxSize int, maxBackup int, maxAge int) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
	}
	return zapcore.AddSync(lumberjackLogger)
}

// GinLogger 替换gin默认logger
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		LG.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("cost", cost.String()),
		)
	}
}
