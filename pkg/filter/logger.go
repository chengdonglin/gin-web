package filter

import (
	"gin-web/pkg/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// GinLogger 替换gin默认logger
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		logs.LG.Info(path,
			zap.String("X-Request-Id", c.Request.Header.Get("X-Request-Id")),
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
