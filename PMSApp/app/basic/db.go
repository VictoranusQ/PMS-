package basic

import (
	"log"

	"PMSApp/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
NewDB 获得数据库连接实例
*/
func NewDB(c *config.Type) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.GetDBConfig().GetURL()), &gorm.Config{})
	if err != nil {
		log.Panicf("[Init] 数据库连接失败: %v", err.Error())
	}

	return db
}
