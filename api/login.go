package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"log"
	"riderStyemProject/global"
	"riderStyemProject/model/request"
	"riderStyemProject/service"
)

func Login(c *gin.Context) {
	var login request.Login
	if err := c.ShouldBind(&login); err != nil {
		log.Println(err)
	}
	data, err := service.Login(&login)
	if  err != nil {
		log.Println(err)
	}
	log.Println(data)
}

func Yzm(c *gin.Context) {
	//生成验证码
	driver := base64Captcha.NewDriverString(60, 240, 0, 0, 4, global.Source, &global.Color, global.Font)
	//把验证码存储到store空间中
	captcha := base64Captcha.NewCaptcha(driver, global.Store)
	//生成验证码的id和图像验证码
	if id, b64s, err := captcha.Generate(); err != nil {
		log.Println(err)
	} else {
		c.JSON(200, gin.H{"id:": id, "b64s:": b64s})
	}
}
