package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"riderStyemProject/global"
	"riderStyemProject/router"
)

func Router() *gin.Engine {
	global.Engine.StaticFS(global.Path, http.Dir(global.Path)) //设置静态文件路径  用来保存上传的文件
	_routerGroup := global.Engine.Group("")
	{
		//登录
		router.Login(_routerGroup)
	}
	routerGroup := global.Engine.Group("")
	routerGroup.Use(router.Jwt()) //jwt鉴权
	{
		router.UserRouter(routerGroup)   //用户路由
		router.RiderRouter(routerGroup)  //骑士路由
		router.OtherRouter(routerGroup)  //其他骑士路由
		router.UploadRouter(routerGroup) //文件上传及其他操作
	}

	return global.Engine
}
