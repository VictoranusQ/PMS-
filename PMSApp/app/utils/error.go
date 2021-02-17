package utils

import "errors"

var (
	ErrBodyCanNotParser = errors.New("请求体解析失败")
	ErrTypeMismatch     = errors.New("查询类型错误")
	ErrParamsMiss       = errors.New("参数不匹配请进行检查")
	ErrPasswordMismatch = errors.New("密码错误")
)
