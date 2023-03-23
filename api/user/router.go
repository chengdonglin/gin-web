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
	handlerUser := New()
	r.GET("/user", handlerUser.GetUser)
}
