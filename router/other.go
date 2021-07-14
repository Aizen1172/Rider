package router

import (
	"github.com/gin-gonic/gin"
	"riderStyemProject/api"
)

func OtherRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("other")
	{
		routerGroup.POST("creatOther", api.CreatOther)           //添加骑士
		routerGroup.PUT("updateOther", api.UpdateOther)          //修改骑士
		routerGroup.DELETE("deleteOther", api.DeleteOther)       //删除骑士
		routerGroup.DELETE("deleteOtherBat", api.DeleteOtherBat) //批量删除骑士
		routerGroup.GET("queryOther", api.QueryOther)            //查询骑士
	}
}
