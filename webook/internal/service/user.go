package service

import (
	"context"
	"gojike/webook/internal/domain"
	"gojike/webook/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository //包私有
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	//考虑加密放在哪里
	hashcode, err := bcrypt.GenerateFromPassword([]byte(u.Passwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Passwd = string(hashcode)
	//然后就是存起来

	return svc.repo.Create(ctx, u)
}
