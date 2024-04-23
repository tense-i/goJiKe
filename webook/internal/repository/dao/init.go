package dao

import "gorm.io/gorm"

//gorm 自动建表
func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &UserInfo{})
}
