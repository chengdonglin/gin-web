package user

import (
	"gin-web/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Route(r *gin.Engine) {
	// 初始化grpc客户端连接
	handlerUser := New()
	r.GET("/user/{id}", handlerUser.GetUser)
}
