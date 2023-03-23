package filter

import (
	"fmt"
	"gin-web/pkg/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印堆栈错误信息
			logs.LG.Error(fmt.Sprintf("panic: %v\n", r))
			debug.PrintStack()
			c.JSON(http.StatusOK, gin.H{
				"code": "-1",
				"msg":  errorToString(r),
				"data": nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
