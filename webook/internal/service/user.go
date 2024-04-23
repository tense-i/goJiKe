package service

import (
	"context"
	"errors"
	"gojike/webook/internal/domain"
	"gojike/webook/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrIncalidUersOrPasswd = errors.New("账户或者密码不对")

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

func (svc *UserService) Login(ctx context.Context, email string, passwd string) (domain.User, error) {
	//先找用户
	u, err := svc.repo.FindByEmail(ctx, email)

	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrIncalidUersOrPasswd
	}
	if err != nil {

		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Passwd), []byte(passwd))
	if err != nil {
		//Debug
		return domain.User{}, ErrIncalidUersOrPasswd
	}
	return u, nil
}

func (svc *UserService) Edit(ctx context.Context, userInfo domain.UserInfo) error {

	return svc.repo.Edit(ctx, userInfo)
}
