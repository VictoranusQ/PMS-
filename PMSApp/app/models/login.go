package models

import (
	"github.com/google/wire"
	"PMSApp/app/daos"
	"PMSApp/app/daos/entity"
	"PMSApp/app/utils"
	"go.uber.org/zap"
	"regexp"
)

var UserSet = wire.NewSet(NewUser, wire.Bind(new(IUser), new(*User)))

type IUser interface {
	Login(username, password string) (error, *UserLoginSchema)
}

type User struct {
	Logger        *zap.Logger
	UserInfoDao   *daos.UserInfo
	UserDetailDao *daos.UserDetail
	Crypto        *utils.Crypto
}

type UserLoginSchema struct {
	UID     string `json:"uid"`
	Manager int    `json:"manager"`
	Admin   int    `json:"admin"`
}

func (u *User) getUserByUID(uid string) (error, *entity.UserInfo) {
	params := daos.UserInfoGetParams{
		UID: uid,
	}
	// 调取数据库数据
	err, temp := u.UserInfoDao.Get(params)
	if err != nil {
		return err, nil
	}

	// 接口转型
	tempInfo, ok := temp.(entity.UserInfo)
	if !ok {
		return utils.ErrTypeMismatch, nil
	}

	return nil, &tempInfo
}

func (u *User) Login(username, password string) (error, *UserLoginSchema) {
	u.Logger.Info("开始进行登录处理", zap.String("username", username))
	var res *entity.UserInfo

	// 识别是否是手机号码
	if len(username) == 4 {
		var err error
		err, res = u.getUserByUID(username)
		if err != nil {
			return err, nil
		}
	} else if match, _ := regexp.MatchString("^1[3-9]\\d{9}", username); match {
		params := daos.UserDetailGetParams{
			Phone: username,
		}
		// 调取数据库数据
		err, temp := u.UserDetailDao.Get(params)
		if err != nil {
			return err, nil
		}

		// 接口转型
		tempInfo, ok := temp.(entity.UserDetail)
		if !ok {
			return utils.ErrTypeMismatch, nil
		}

		// 查找用户
		err, res = u.getUserByUID(tempInfo.HrID)
		if err != nil {
			return err, nil
		}
		res.Detail = tempInfo
	} else {
		return utils.ErrParamsMiss, nil
	}

	temp := u.Crypto.MD5(password)
	if temp != res.Password {
		u.Logger.Debug("密码打印",
			zap.String("name", res.Detail.Phone),
			zap.String("temp", temp),
			zap.String("password", res.Password),
		)
		return utils.ErrPasswordMismatch, nil
	}

	var isManager int
	if res.IsLeader == 1 || res.IsHr == 1 {
		isManager = 1
	} else {
		isManager = 0
	}

	return nil, &UserLoginSchema{
		UID:     res.UID,
		Manager: isManager,
		Admin:   res.IsAdmin,
	}
}

func NewUser(userDao *daos.UserInfo, userDetail *daos.UserDetail, logger *zap.Logger) *User {
	return &User{
		UserInfoDao:   userDao,
		UserDetailDao: userDetail,
		Logger:        logger,
	}
}
