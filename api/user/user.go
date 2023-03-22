package user

import (
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "success")
	return
}
