package middleware

import (
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine, filters ...gin.HandlerFunc) *gin.Engine {
	r.Use(filters...)
	return r
}
