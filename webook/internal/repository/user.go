package repository

import (
	"context"
	"gojike/webook/internal/domain"
	"gojike/webook/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

// 邮箱重复错误
var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
var ErrUserNotFound = dao.ErrUserNotFound

// 要用的东西不要自己初始化--传参进来
func NewUsrRepostory(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:     u.Id,
		Email:  u.Email,
		Passwd: u.Passwd,
	}, nil

}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:  u.Email,
		Passwd: u.Passwd,
		//时间的存放问题
	})

	//在这里操作缓存
}
func (r *UserRepository) FindById(int64) {
	//先往cache里面找
	//再从dao里面找
	//找到了回写cache
}

func (r *UserRepository) Edit(ctx context.Context, userinfo domain.UserInfo) error {

	//return r.dao.InsertUserInfo(ctx, dao.UserInfo{
	//	NickName: userinfo.NickName,
	//	//PhoneNum: userinfo.PhoneNum,
	//	Birthday: userinfo.Birthday,
	//	Aboutme:  userinfo.Aboutme,
	//})
	return nil
}
