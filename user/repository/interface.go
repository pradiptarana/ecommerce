package repository

import (
	model "github.com/pradiptarana/user/model/user"
)

//go:generate mockgen -destination=../mocks/mock_user.go -package=mocks github.com/pradiptarana/user/repository UserRepository
type UserRepository interface {
	SignUp(us *model.User) error
	GetUser(username string) (*model.User, error)
}
