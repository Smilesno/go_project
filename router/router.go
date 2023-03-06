package router

import (
	"github.com/gin-gonic/gin"
	swaggofiles "github.com/swaggo/files"
	ginSwaggo "github.com/swaggo/gin-swagger"
	"project/docs"
	"project/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwaggo.WrapHandler(swaggofiles.Handler))
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/user", service.GetUserList)
	return r
}
