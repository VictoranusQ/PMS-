package entity

import "gorm.io/gorm"

// UserDetail user_detail 表模型
type UserDetail struct {
	ID        int64  `gorm:"column:id;type=int;not null;primaryKey;autoIncrement"`
	HrID      string `gorm:"column:hrid;type=varchar(255);not null"`
	Name      string `gorm:"type=varchar(255);not null"`
	Sex       int    `gorm:"type=tinyint;not null"`
	Stuid     string `gorm:"type=varchar(255);not null"`
	College   string `gorm:"type=varchar(255);not null"`
	Major     string `gorm:"type=varchar(255);not null"`
	Class     string `gorm:"type=varchar(255);not null"`
	Dormitory string `gorm:"type=varchar(255);not null"`
	Phone     string `gorm:"type=varchar(255);not null"`
	QQ        string `gorm:"column:qq;type=varchar(255);not null"`
	Email     string `gorm:"type=varchar(255);not null"`
	Grade     string `gorm:"type=varchar(255);not null"`
}

func (UserDetail) TableName() string {
	return "user_detail"
}

// GetUserDetail 获得 user_detail 模型
func GetUserDetail(db *gorm.DB) *gorm.DB {
	return db.Model(new(UserDetail))
}
