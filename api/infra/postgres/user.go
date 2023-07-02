package postgres

import (
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