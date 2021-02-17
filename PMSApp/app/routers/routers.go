package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"PMSApp/app/controllers"
)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine)
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	LoginAPI *controllers.Login
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) {
	a.RegisterAPI(app)
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api",
	}
}
