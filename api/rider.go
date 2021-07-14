package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
	"riderStyemProject/service"
)

//CreatRider 添加骑士
func CreatRider(c *gin.Context) {
	var rider request.CreatRider
	if err := c.ShouldBindJSON(&rider); err != nil {
		log.Println(err)
	}
	if err := service.CreatRider(&rider); err != nil {
		log.Println(err)
	}
}

//UpdateRider 修改骑士
func UpdateRider(c *gin.Context) {
	var rider request.UpdateRider
	if err := c.ShouldBindJSON(&rider); err != nil {
		log.Println(err)
	}
	if err := service.UpdateRider(&rider); err != nil {
		log.Println(err)
	}
}

//DeleteRider 删除骑士
func DeleteRider(c *gin.Context) {
	var rider request.DeleteRider
	if err := c.ShouldBindJSON(&rider); err != nil {
		log.Println(err)
	}
	if err := service.DeleteRider(&rider); err != nil {
		log.Println(err)
	}
}

//DeleteRiderBat 批量删除骑士
func DeleteRiderBat(c *gin.Context) {
	var riderId global.Ids
	if err := c.ShouldBindJSON(&riderId); err != nil {
		log.Println(err)
	}
	if err := service.DeleteRiderBat(&riderId); err != nil {
		log.Println(err)
	}
}

//QueryRider 查询骑士
func QueryRider(c *gin.Context) {
	var rider model.Rider
	if err := c.ShouldBindJSON(&rider); err != nil {
		log.Println(err)
	}
	if err, _rider := service.QueryRider(&rider); err != nil {
		log.Println(err)
	} else {
		c.JSON(200, gin.H{"rider": _rider})
	}
}
