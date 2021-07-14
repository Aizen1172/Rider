package router

import (
	"github.com/gin-gonic/gin"
	"riderStyemProject/api"
)

func UploadRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("files")
	{
		routerGroup.POST("upload",api.Upload)//文件上传
		routerGroup.DELETE("delete",api.Delete) //文件删除
	}
}
