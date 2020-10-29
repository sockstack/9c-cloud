package repository

import "github.com/sockstack/9c-cloud/auth/model"

type UserRepository struct {
	user *model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{user: new(model.User)}
}

func (u *UserRepository) FindUserByUsername(username string) (user model.User, err error) {
	return
}
