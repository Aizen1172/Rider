package service

import (
	"errors"
	"gorm.io/gorm"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
)

//CreatRider 添加骑士
func CreatRider(rider *request.CreatRider) (err error) {
	if errors.Is(global.DB.Where("name = ?", rider.Name).First(&model.Rider{}).Error, gorm.ErrRecordNotFound) {
		_rider := rider.Creat()
		err = global.DB.Create(&_rider).Error
	} else {
		return errors.New("骑士已存在，请重新命名。")
	}
	return err
}

//UpdateRider 修改骑士
func UpdateRider(rider *request.UpdateRider) (err error) {
	_rider := rider.Update()
	if errors.Is(global.DB.Where("id = ? AND name = ?", rider.ID, rider.Name).First(&model.Rider{}).Error, gorm.ErrRecordNotFound) {
		if errors.Is(global.DB.Where("name = ?", rider.Name).First(&model.Rider{}).Error, gorm.ErrRecordNotFound) {
			return global.DB.Where("id = ?", rider.ID).Updates(&_rider).Error
		} else {
			return errors.New("骑士已存在，请重新命名。")
		}
	}
	return global.DB.Where("id = ?", rider.ID).Updates(&_rider).Error
}

//DeleteRider 删除骑士
func DeleteRider(rider *request.DeleteRider) (err error) {
	//根据id删除当前表的信息
	/*err = global.DB.Where("id = ?", rider.ID).Delete(&model.Rider{}).Error*/
	//删除全部信息
	/*var _rider []model.Rider
	err = global.DB.Preload("Others").Find(&_rider).Error
	err = global.DB.Select("Others").Delete(&_rider).Error*/
	//根据id删除主表和子表的信息
	var _rider []model.Rider
	err = global.DB.Preload("Others").Where("id = ?", rider.ID).Find(&_rider).Error
	err = global.DB.Select("Others").Where("id = ?", rider.ID).Delete(&_rider).Error
	return err
}

//DeleteRiderBat 批量删除骑士
func DeleteRiderBat(rider *global.Ids) (err error) {
	//根据id删除当前表的信息(不包括子信息)
	/*err = global.DB.Where("id in ?", rider.Ids).Delete(&model.Rider{}).Error*/
	//根据id删除主表和子表的信息
	var _rider []model.Rider
	err = global.DB.Preload("Others").Where("id in ?", rider.Ids).Find(&_rider).Error
	err = global.DB.Select("Others").Where("id in ?", rider.Ids).Delete(&_rider).Error
	return err
}

//QueryRider 查询骑士
func QueryRider(rider *model.Rider) (err error, _rider []model.Rider) {
	//查询全部信息
	/*err = global.DB.Preload("Others").Find(&_rider).Error*/
	//根据id查询当前表的信息(不包括子信息)
	/*err = global.DB.Where("id = ?", rider.ID).First(&_rider).Error*/
	//根据id或年份查询主表和子表信息
	err = global.DB.Preload("Others").Where("id = ? OR year =?", rider.ID, rider.Year).Find(&_rider).Error 
	return err, _rider
}
