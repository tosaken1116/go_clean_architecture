package service

import (
	"go-clean-archtecture/api/domain/model"
	"go-clean-archtecture/api/domain/repository"
)

type IUserService interface{
	CreateNewUser(u model.User)(model.User,error)
}

type userService struct {
	repo repository.IUserRepository
}

func NewStudentService(ur repository.IUserRepository) IUserService{
	return &userService{
		repo: ur,
	}
}

func (us *userService) CreateNewUser(u model.User) (model.User,error){
	return us.repo.CreateNewUser(u)
}