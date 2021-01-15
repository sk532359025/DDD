package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sk532359025/DDD/adapter/controller"
)


func Router() *gin.Engine{
	router := gin.Default()
	router.POST("/login", controller.UserLogin)
	return router
}