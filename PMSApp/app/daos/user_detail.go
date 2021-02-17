package daos

import (
	"github.com/google/wire"
	"PMSApp/app/daos/entity"
	"PMSApp/app/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UserDetailSet User DI
var UserDetailSet = wire.NewSet(wire.Struct(new(UserDetail), "*"))

// UserDetail 用户相关
type UserDetail struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

type UserDetailGetParams struct {
	Phone string
}

// Get 根据条件获取信息
func (u *UserDetail) Get(params interface{}) (error, interface{}) {
	resultParams, ok := params.(UserDetailGetParams)
	if !ok {
		u.Logger.Error("参数传递错误, 应为 UserDetailGetParams 类型")
		return utils.ErrTypeMismatch, nil
	}

	db := entity.GetUserDetail(u.DB)

	if v := resultParams.Phone; v != "" {
		u.Logger.Info("使用 phone 查找", zap.String("uid", v))
		db = db.Where(map[string]interface{}{"phone": v})
	}

	var user entity.UserDetail
	if err := db.First(&user).Error; err != nil {
		return err, nil
	}

	return nil, user
}
