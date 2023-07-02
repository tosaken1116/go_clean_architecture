package router

import (
	"go-clean-architecture/api/adapter/http/handler"
	"go-clean-architecture/api/domain/service"
	"go-clean-architecture/api/infra"
	"go-clean-architecture/api/infra/postgres"
	"go-clean-architecture/api/usecase"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(ur *gin.RouterGroup){
	postgresConn := infra.NewPostgresConnector()
	userRepository := postgres.NewUserRepository(postgresConn.Conn)
	userService := service.NewStudentService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userHandler := handler.NewUserHandler(userUsecase)
	ur.POST("/create",userHandler.CreateNewUser())
	ur.GET("/check_screen_id/:screen_id",userHandler.CheckScreenId())
}