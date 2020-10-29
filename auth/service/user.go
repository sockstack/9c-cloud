package service

import (
	"github.com/sockstack/9c-cloud/auth/contract/repository"
	"github.com/sockstack/9c-cloud/auth/model"
)

type UserService struct {
	mapper repository.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Login(dto model.UserDto) (userDto model.UserDto, err error) {
	user, err := u.mapper.FindUserByUsername(dto.Username)

	userDto.ID = user.ID
	userDto.Username = user.Username
	userDto.Password = user.Password
	return
}

func (u *UserService) Unwrap(data interface{}, err error) *UserService {
	if err != nil {
		panic(err)
	}
	return u
}
