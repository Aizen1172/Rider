package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
	"riderStyemProject/model/response"
)

func Login(login *request.Login) (result *response.Login,err error) {
	var user model.User
	if global.DB.Where("name = ?", login.Name).First(&user).Error != nil {
		return nil, errors.New("用户名不正确或不存在。")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)) != nil {
		return nil, errors.New("密码错误！")
	}
	//对验证码进行验证   verify
	if login.YzmId == "" || login.Yzm == "" {
		return nil, errors.New("验证码不能为空! ")
	} else if global.Store.Verify(login.YzmId, login.Yzm, true) == false {
		return nil, errors.New("验证码错误。")
	}
		return Token(&user)
}
