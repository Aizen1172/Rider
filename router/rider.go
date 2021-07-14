package router

import (
	"github.com/gin-gonic/gin"
	"riderStyemProject/api"
)

func RiderRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("rider")
	{
		routerGroup.POST("creatRider", api.CreatRider)           //添加骑士
		routerGroup.PUT("updateRider", api.UpdateRider)          //修改骑士
		routerGroup.DELETE("deleteRider", api.DeleteRider)       //删除骑士
		routerGroup.DELETE("deleteRiderBat", api.DeleteRiderBat) //批量删除骑士
		routerGroup.GET("queryRider", api.QueryRider)            //查询骑士
	}

}
