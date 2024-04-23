package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB //包私有
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

// 传结构体而不是指针--大概率会分配到栈上、而且本身数据不大
func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now

	//withContxt传递数据库的上下文、保持db的统一性
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConfictsErrno uint16 = 1062
		//重复邮箱冲突
		if mysqlErr.Number == uniqueConfictsErrno {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) InsertUserInfo(ctx context.Context, u UserInfo) error {

	err := dao.db.WithContext(ctx).Create(&u).Error
	// if mysqlErr, ok := err.(*mysql.MySQLError); ok {
	// 	const uniqueConfictsErrno uint16 = 1062
	// 	//重复邮箱冲突
	// 	if mysqlErr.Number == uniqueConfictsErrno {
	// 		return ErrUserDuplicateEmail
	// 	}
	// }
	if err != nil {
		print(err)
	}
	return nil
}

// Use 直接对用数据库表结构
type User struct {
	Id     int64  `gorm:"primaryKey,autoIncrement"`
	Email  string `gorm:"unique"`
	Passwd string
	Utime  int64 //更新时间
	Ctime  int64 //创建时间
}

type UserInfo struct {
	NickName string
	//Email    string `gorm:"unique"`
	//PhoneNum string `gorm:"primaryKey"`
	Birthday string
	Aboutme  string
}
