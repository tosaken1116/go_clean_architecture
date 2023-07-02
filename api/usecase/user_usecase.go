package usecase

import (
	"go-clean-archtecture/api/domain/service"
	usecase_model "go-clean-archtecture/api/usecase/model"
)

type IUserUsecase interface{
	CreateNewUser(u usecase_model.PostUser) (*usecase_model.User,error)
}

type userUsecase struct{
	svc service.IUserService
}

func NewUserUsecase(us service.IUserService) IUserUsecase {
	return &userUsecase{
		svc: us,
	}
}

func (uu *userUsecase) CreateNewUser(postUser usecase_model.PostUser) (*usecase_model.User,error){
	newUser := usecase_model.UserFromPostModel(&postUser)
	u,err := uu.svc.CreateNewUser(newUser)
	if err != nil{
		return nil,err
	}
	user:= usecase_model.UserFromDomainModel(&u)
	return user,nil
}