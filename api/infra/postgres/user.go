package postgres

import (
	"errors"
	"fmt"
	"go-clean-archtecture/api/domain/model"
	"go-clean-archtecture/api/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository{
	return &userRepository{
		DB:db,
	}
}

func (ur *userRepository) CreateNewUser(u model.User) (user model.User, err error) {
	if err = ur.DB.Create(&u).Error;err != nil {
		return
	}
	user = u
	return
}
func (ur *userRepository) GetAllScreenId(screenId string)(isDuplicate bool,err error){
	duplicateScreenIdUser := new(model.User)
	if err = ur.DB.Where("screen_id = ?",screenId).First(&duplicateScreenIdUser).Error; err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			isDuplicate = true
			err = nil
		}
		return
	}
	isDuplicate = false
	return
}