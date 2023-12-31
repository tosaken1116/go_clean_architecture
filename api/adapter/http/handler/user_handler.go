package handler

import (
	"go-clean-architecture/api/usecase"
	"go-clean-architecture/api/usecase/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct{
	usecase usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *userHandler{
	return &userHandler{
		usecase: uu,
	}
}

func (uh *userHandler) CreateNewUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		newUser:= new(model.PostUser)
		if err := c.ShouldBindJSON(newUser); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "parse failed",
			})
			return
		}

		user,err := uh.usecase.CreateNewUser(*newUser)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "create failed",
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"data":user,
			})
		}
	}
}

func (uh *userHandler) CheckScreenId() gin.HandlerFunc{
	return func(c *gin.Context) {
		isDuplicate,err := uh.usecase.CheckScreenId(c.Param("screen_id"))
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "check failed",
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"data":isDuplicate,
			})
		}
	}
}