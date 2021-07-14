package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
	"riderStyemProject/service"
)

//CreatUser  添加用户
func CreatUser(c *gin.Context) {
	var user request.CreatUser
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
	}
	if err := service.CreatUser(&user); err != nil {
		log.Println(err)
	}
}

// UpdateUser 修改用户
func UpdateUser(c *gin.Context) {
	var user request.UpdateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
	}

	if err := service.UpdateUser(&user); err != nil {
		log.Println(err)
	}
}

// UpdatePassword 修改用户密码
func UpdatePassword(c *gin.Context) {
	var user request.UpdateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
	}

	if err := service.UpdatePassword(&user); err != nil {
		log.Println(err)
	}
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	var user request.DeleteUser
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
	}

	if err := service.DeleteUser(&user); err != nil {
		log.Println(err)
	}
}

// DeleteUserBat 批量删除用户
func DeleteUserBat(c *gin.Context) {
	var userId global.Ids
	if err := c.ShouldBindJSON(&userId); err != nil {
		log.Println(err)
	}

	if err := service.DeleteUserBat(userId); err != nil {
		log.Println(err)
	}
}

// QueryUser 查询用户
func QueryUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
	}

	if err, _user := service.QueryUser(&user); err != nil {
		log.Println(err)
	} else {
		c.JSON(200, gin.H{"user": _user})
	}
}
