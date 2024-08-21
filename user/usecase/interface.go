package usecase

import (
	model "github.com/pradiptarana/user/model/user"
)

type UsersUsecase interface {
	SignUp(req *model.User) error
	Login(req *model.LoginRequest) (string, error)
}
