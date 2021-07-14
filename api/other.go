package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
	"riderStyemProject/service"
)

//CreatOther 添加骑士
func CreatOther(c *gin.Context) {
	var other request.CreatOther
	if err := c.ShouldBindJSON(&other); err != nil {
		log.Println(err)
	}
	if err := service.CreatOther(&other); err != nil {
		log.Println(err)
	}
}

//UpdateOther 修改骑士
func UpdateOther(c *gin.Context) {
	var other request.UpdateOther
	if err := c.ShouldBindJSON(&other); err != nil {
		log.Println(err)
	}
	if err := service.UpdateOther(&other); err != nil {
		log.Println(err)
	}
}

//DeleteOther 删除骑士
func DeleteOther(c *gin.Context) {
	var other request.DeleteOther
	if err := c.ShouldBindJSON(&other); err != nil {
		log.Println(err)
	}
	if err := service.DeleteOther(&other); err != nil {
		log.Println(err)
	}
}

//DeleteOtherBat 批量删除骑士
func DeleteOtherBat(c *gin.Context) {
	var otherId global.Ids
	if err := c.ShouldBindJSON(&otherId); err != nil {
		log.Println(err)
	}
	if err := service.DeleteOtherBat(&otherId); err != nil {
		log.Println(err)
	}
}

//QueryOther 查询骑士
func QueryOther(c *gin.Context) {
	var other model.Other
	if err := c.ShouldBindJSON(&other); err != nil {
		log.Println(err)
	}
	if err, _other := service.QueryOther(&other); err != nil {
		log.Println(err)
	} else {
		c.JSON(200, gin.H{"Other": _other})
	}
}

