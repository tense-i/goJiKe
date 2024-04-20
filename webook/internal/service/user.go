package service

import (
	"context"
	"gojike/webook/internal/domain"
	"gojike/webook/internal/repository"
	"gojike/webook/internal/repository/dao"
)

type UserService struct {
	repo *repository.UserRepository //包私有
	dao  *dao.UserDAO
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	//考虑加密放在哪里
	//然后就是存起来

	return svc.repo.Create(ctx, u)
}
