package model

import (
	"go-clean-archtecture/api/domain/model"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string  `json:"id"`
	ScreenID  string `json:"screen_id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	IsPublic  bool   `json:"is_public"`
}

type PostUser struct {
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	ScreenID string `json:"screen_id"`
}

func UserFromDomainModel(m *model.User) *User{
	u := &User{
		ID: m.ID,
		ScreenID: m.ScreenID,
		UserName: m.UserName,
		Email: m.Email,
		IsPublic: m.IsPublic,
	}
	return u
}

func UserFromPostModel(m *PostUser) model.User{
	hashedPassword,_ :=bcrypt.GenerateFromPassword([]byte(m.Password),10)
	u := &model.User{
		ID: uuid.New().String(),
		ScreenID: m.ScreenID,
		PasswordHash: string(hashedPassword),
		UserName: m.UserName,
		IsPublic: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return *u
}