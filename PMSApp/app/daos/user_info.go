package daos

import (
	"github.com/google/wire"
	"PMSApp/app/daos/entity"
	"PMSApp/app/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UserInfoSet User DI
var UserInfoSet = wire.NewSet(wire.Struct(new(UserInfo), "*"))

// UserInfo 用户相关
type UserInfo struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

type UserInfoGetParams struct {
	UID string
}

// Get 根据条件获取信息
func (u *UserInfo) Get(params interface{}) (error, interface{}) {
	resultParams, ok := params.(UserInfoGetParams)
	if !ok {
		u.Logger.Error("参数传递错误, 应为 UserInfoGetParams 类型")
		return utils.ErrTypeMismatch, nil
	}

	db := entity.GetUserInfo(u.DB)

	if v := resultParams.UID; v != "" {
		u.Logger.Info("使用 uid 查找", zap.String("uid", v))
		db = db.Where(map[string]interface{}{"uid": v})
	}

	var user entity.UserInfo
	if err := db.First(&user).Error; err != nil {
		return err, nil
	}

	return nil, user
}
