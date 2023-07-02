package service

import (
	"fmt"
	"go-clean-archtecture/api/domain/model"
	"go-clean-archtecture/api/domain/repository"
)

type IUserService interface{
	CreateNewUser(u model.User)(model.User,error)
	CheckScreenId(screenId string)(*bool, error)
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
func (us *userService) CheckScreenId(screenId string) (*bool, error){
	isDuplicate,err := us.repo.GetAllScreenId(screenId)
	if err != nil{
		return nil,err
	}
	return &isDuplicate,nil
}