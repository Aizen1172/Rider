package router

import (
	"github.com/gin-gonic/gin"
	"riderStyemProject/api"
)

func Login(router *gin.RouterGroup)  {
	routerGroup := router.Group("login")
	{
		routerGroup.POST("login",api.Login)
		routerGroup.POST("yzm",api.Yzm)
	}

}
