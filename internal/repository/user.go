package repository

import (
	"context"
	"gojike/webook/internal/domain"
	"gojike/webook/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

// 要用的东西不要自己初始化--传参进来
func NewUsrRepostory(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
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
