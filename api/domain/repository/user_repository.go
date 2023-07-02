package repository

import (
	"go-clean-archtecture/api/domain/model"
)

type IUserRepository interface {
	CreateNewUser(u model.User) (model.User, error)
}

