package service

import "github.com/sockstack/9c-cloud/auth/model"

type IAuthService interface {
	Login(dto model.UserDto) (userDto model.UserDto, err error)
}
