package router

import (
	"audition/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/javaAuditionBefor", handler.GetMultiple)
	router.POST("/register", handler.Registered)
	router.POST("/login", handler.Login)
	return router
}
