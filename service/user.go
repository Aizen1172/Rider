package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
)

//CreatUser 添加用户
func CreatUser(user *request.CreatUser) (err error) {
	if errors.Is(global.DB.Where("name = ?", user.Name).First(&model.User{}).Error,gorm.ErrRecordNotFound){
		if len(user.Password) < 6 {
			return errors.New("密码长度必须大于等于6。 ")
		}
		user.Password = model.GenerateFromPassword(user.Password)
		_user := user.Creat()
		err = global.DB.Create(&_user).Error
	} else {
		return errors.New("用户名已存在。 ")
	}
	return err
}

// UpdateUser 修改角色
func UpdateUser(user *request.UpdateUser) (err error) {
	if errors.Is(global.DB.Where("name = ?", user.Name).First(&model.User{}).Error,gorm.ErrRecordNotFound){
		_user := user.Update()
		err = global.DB.Where("id = ?",user.ID).Updates(&_user).Error
	} else {
		return errors.New("用户名重复。")
	}
	return err
}

// UpdatePassword 修改角色密码
func UpdatePassword(user *request.UpdateUser) (err error) {
	var _user = model.User{}
	if errors.Is(global.DB.Where("id = ? AND name = ?", user.ID,user.Name).First(&_user).Error,gorm.ErrRecordNotFound) {
		return errors.New("用户名错误或者不存在。")
	}
	//密码校验
	if bcrypt.CompareHashAndPassword([]byte(_user.Password), []byte(user.Password)) != nil {
		_user.Password = user.Password
		if len(_user.Password) < 6 {
			return errors.New("密码长度必须大于等于6。")
		}
		_user.Password = model.GenerateFromPassword(_user.Password)
		err = global.DB.Where("id = ?",user.ID).Updates(&_user).Error
	} else {
		return errors.New("密码一致。")
	}
	return err
}

// DeleteUser 删除角色
func DeleteUser(user *request.DeleteUser) (err error) {
	err = global.DB.Where("id = ?",user.ID).Delete(&model.User{}).Error
	return err
}

// DeleteUserBat 批量删除角色
func DeleteUserBat(userId global.Ids) (err error) {
	err = global.DB.Delete(&model.User{}, "id in ?", userId.Ids).Error
	return err
}

// QueryUser 查询角色
func QueryUser(user *model.User) (err error, _user model.User) {
	err = global.DB.Where("id = ?", user.ID).Find(&_user).Error
	return
}
