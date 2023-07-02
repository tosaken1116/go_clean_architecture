package usecase

import (
	"go-clean-archtecture/api/domain/service"
	usecase_model "go-clean-archtecture/api/usecase/model"
)

type IUserUsecase interface{
	CreateNewUser(u usecase_model.PostUser) (*usecase_model.User,error)
	CheckScreenId(screen_id string) (*bool,error)
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
func (uu *userUsecase) CheckScreenId(screenId string) (*bool,error){
	isDuplicate,err := uu.svc.CheckScreenId(screenId)
	if err != nil{
		return nil,err
	}
	return isDuplicate,nil
}