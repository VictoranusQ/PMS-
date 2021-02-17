package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"PMSApp/app/models"
	"PMSApp/app/utils"
	"go.uber.org/zap"
)

// LoginSet Login DI
var LoginSet = wire.NewSet(wire.Struct(new(Login), "*"))

// Login 登录结构体
type Login struct {
	Logger    *zap.Logger
	Util      utils.Ginx
	UserModel models.IUser
}

// 登录结构体
type loginSchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 登录方法
func (l *Login) Login(c *gin.Context) {
	var data loginSchema
	if err := l.Util.ParseJSON(c, &data); err != nil {
		l.Util.FailWarp(c, utils.ErrBodyCanNotParser)
		return
	}

	err, userInfo := l.UserModel.Login(data.Username, data.Password)
	if err != nil {
		l.Logger.Warn("服务处理失败", zap.Error(err))
		l.Util.FailWarp(c, err)
		return
	}

	l.Util.SuccessWarp(c, userInfo)
}
