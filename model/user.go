package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"comment:用户名"`
	Password string `json:"password" form:"password" gorm:"comment:密码"`
}

// GenerateFromPassword 密码加密  bcrypt
func GenerateFromPassword(password string) string {
	if psw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		fmt.Println(err)
	} else {
		password = string(psw)
	}
	return password
}
