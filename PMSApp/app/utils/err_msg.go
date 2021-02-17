package utils

import "gorm.io/gorm"

type returnError struct {
	code string
	msg  string
}

var errorMap = map[error]returnError{
	ErrBodyCanNotParser: {
		code: "1001",
		msg:  "Request Body 解析失败",
	},
	ErrParamsMiss: {
		code: "1002",
		msg:  "参数错误，请检查",
	},
	gorm.ErrRecordNotFound: {
		code: "1011",
		msg:  "记录未找到",
	},
	ErrTypeMismatch: {
		code: "1021",
		msg:  "服务转型失败",
	},
	ErrPasswordMismatch: {
		code: "1031",
		msg:  "用户密码错误",
	},
}
