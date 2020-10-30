package service

import "github.com/sockstack/9c-cloud/auth/model"

type IUserService interface {
	Login(dto model.UserDto) (userDto model.UserDto, err error)
}
