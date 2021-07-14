package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"riderStyemProject/global"
	"riderStyemProject/service"
)

func Upload(c *gin.Context) {
	if form, err := c.MultipartForm();err != nil {
		log.Println(err)
	}else {
		files := form.File["file"]
		for _, file := range files {
			if err := service.Upload(file);err != nil {
				log.Println(err)
			}
		}
	}
}

func Delete(c *gin.Context) {
	fileId := &global.Ids{}
	if err := c.ShouldBindJSON(fileId);err != nil {
		log.Println(err)
	}
	if err := service.Delete(fileId.Ids	);err != nil {
		log.Println(err)
	}
}
