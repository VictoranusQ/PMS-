package routers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// RegisterAPI 路由列表
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group(strings.Join(a.Prefixes(), ""))
	{
		g.POST("login", a.LoginAPI.Login)
	}
}
