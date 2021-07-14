package main

import (
	"fmt"
	"riderStyemProject/global"
	"riderStyemProject/initialize"
	"riderStyemProject/logger"
)

func main() {
	logger.Logger()
	initialize.Gorm()   //初始化数据库
	initialize.Router() //初始化路由

	if err := global.Engine.Run(":8080"); err != nil {
		fmt.Println("服务启动失败。",err)
		return
	}
}
