package router

import "github.com/gin-gonic/gin"


const (
	apiVersion ="/v1"
	userApiRoot =  "/users"
)

func InitRouter() *gin.Engine{
	r := gin.Default()
	api := r.Group("api")
	{
		v1:=api.Group(apiVersion)
		{
			ur := v1.Group(userApiRoot)
			InitUserRouter(ur)
		}
	}
	return r
}