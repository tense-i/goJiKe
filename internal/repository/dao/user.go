package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB //包私有
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// 传结构体而不是指针--大概率会分配到栈上、而且本身数据不大
func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now

	//withContxt传递数据库的上下文、保持db的统一性
	return dao.db.WithContext(ctx).Create(&u).Error

}

// Use 直接对用数据库表结构
type User struct {
	Id     int64  `gorm:"primaryKey,autoIncrement"`
	Email  string `gorm:"unique"`
	Passwd string
	Utime  int64 //更新时间
	Ctime  int64 //创建时间
}
