package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
	"image/color"
	"os"
)

var (
	DB     *gorm.DB
	Engine = gin.Default()
	Path   = "files/img"
	Source = "1234567890qwertyuioplkjhgfdsazxcvbnm"
	Color  = color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}
	Font   []string
	Store  = base64Captcha.DefaultMemStore //创建验证码存储空间
	Secret = []byte("Aizen")
)

type Ids struct {
	Ids []int `json:"ids"`
}

// IsExist 判断文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path) //返回该路径的文件信息
	if err != nil {
		//判断err错误报告文件是否存在
		if os.IsExist(err) {
			return true
		}
		//判断err错误报告文件是否不存在
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
