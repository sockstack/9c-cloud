package model

type UserDto struct {
	ID uint32
	Username string
	Password string
}

type User struct {
	ID uint32
	Username string
	Password string
}

type UserQuery struct {
	Username string `json:"username" form:"username" binding:"require"`
	Password string `json:"password" form:"password" binding:"require"`
}

func NewUserQuery() *UserQuery {
	return &UserQuery{}
}
