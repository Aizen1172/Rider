package router

import (
	"github.com/gin-gonic/gin"
	"riderStyemProject/api"
)

func UserRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("user")
	{
		routerGroup.POST("creatUser", api.CreatUser)           //新增用户
		routerGroup.PUT("updateUser", api.UpdateUser)          //修改用户
		routerGroup.PUT("updatePassword", api.UpdatePassword)  //修改用户密码
		routerGroup.DELETE("deleteUser", api.DeleteUser)       //删除用户
		routerGroup.DELETE("deleteUserBat", api.DeleteUserBat) //批量删除用户
		routerGroup.GET("queryUser", api.QueryUser)            //查询用户
	}
}
