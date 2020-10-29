package repository

import "github.com/sockstack/9c-cloud/auth/model"

type IUserRepository interface {
	FindUserByUsername(username string) (user model.User, err error)
}
