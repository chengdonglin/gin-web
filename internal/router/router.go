package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router 路由接口
type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	// 传入Router接口的实现
	//rg := New()
	//rg.Route(&user.RouterUser{}, r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "The incorrect API route")
	})
	for _, ro := range routers {
		ro.Route(r)
	}
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
