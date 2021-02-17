package daos

import "github.com/google/wire"

// OrderOpts 排序基础结构
type OrderOpts struct {
	Name string
	Desc bool
}

// PaginationOpts 分页参数
type PaginationOpts struct {
	OnlyCount bool
	Page      int
	limit     int
}

// Opts 查询选项
type Opts struct {
	Order *[]OrderOpts
	Page  *PaginationOpts
}

// IDaos daos 层需要实现的接口
type IDaos interface {
	// 本方法应为只从数据库中 select 一条信息
	Get(params interface{}) (error, interface{})
}

// DaoSet daos DI
var DaoSet = wire.NewSet(
	UserInfoSet,
	UserDetailSet,
)
