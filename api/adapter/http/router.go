package http

import "github.com/gin-gonic/gin"


const (
	apiVersion ="/v1"
)

func initRouter() *gin.Engine{
	r := gin.Default()
	return r
}