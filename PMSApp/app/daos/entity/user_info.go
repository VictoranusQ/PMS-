package entity

import "gorm.io/gorm"

// UserInfo user_info 表模型
type UserInfo struct {
	ID       int64      `gorm:"column:id;type=int;not null;primaryKey;autoIncrement"`
	UID      string     `gorm:"column:uid;type=varchar(255);not null"`
	Password string     `gorm:"type:varchar(255);not null"`
	Role     int        `gorm:"type:tinyint;not null"`
	Status   string     `gorm:"type:varchar(255);default:'0';not null"`
	IsLeader int        `gorm:"type:tinyint(2);default:0;not null"`
	IsHr     int        `gorm:"type:tinyint(2);default:0;not null"`
	IsAdmin  int        `gorm:"type:tinyint(2);default:0;not null"`
	Time     int        `gorm:"type:int;"`
	IsDelete int        `gorm:"type:tinyint(2);default:0;not null"`
	Detail   UserDetail `gorm:"foreignKey:HrID;references:uid"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

// GetUserInfo 获得 user_info 模型
func GetUserInfo(db *gorm.DB) *gorm.DB {
	return db.Model(new(UserInfo))
}
