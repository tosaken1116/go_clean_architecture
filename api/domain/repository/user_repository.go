package repository

import (
	"go-clean-architecture/api/domain/model"
)

type IUserRepository interface {
	CreateNewUser(u model.User) (model.User, error)
	GetAllScreenId(screenId string)(bool, error)
}

