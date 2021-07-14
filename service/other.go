package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
	"time"
)

//CreatOther 添加骑士
func CreatOther(other *request.CreatOther) (err error) {
	if errors.Is(global.DB.Where("name = ?", other.Name).First(&model.Other{}).Error, gorm.ErrRecordNotFound) {
		var _other = other.Creat()
		err = global.DB.Create(&_other).Error
	} else {
		return errors.New("骑士已存在，请重新命名。")
	}
	return err
}

//UpdateOther 修改骑士
func UpdateOther(other *request.UpdateOther) (err error) {
	_other := other.Update()
	if errors.Is(global.DB.Where("id = ? AND name = ?", other.ID, other.Name).First(&model.Other{}).Error, gorm.ErrRecordNotFound) {
		if errors.Is(global.DB.Where("name = ?", other.Name).First(&model.Other{}).Error, gorm.ErrRecordNotFound) {
			return global.DB.Where("id = ?", other.ID).Updates(&_other).Error
		} else {
			return errors.New("骑士已存在，请重新命名。")
		}
	}
	return global.DB.Where("id = ?", other.ID).Updates(&_other).Error
}

//DeleteOther 删除骑士
func DeleteOther(other *request.DeleteOther) (err error) {
	err = global.DB.Where("id = ?", other.ID).Delete(&model.Other{}).Error
	return err
}

//DeleteOtherBat 批量删除骑士
func DeleteOtherBat(other *global.Ids) (err error) {
	err = global.DB.Where("id in ?", other.Ids).Delete(&model.Other{}).Error
	return err
}

//QueryOther 查询骑士
func QueryOther(other *model.Other) (err error, _other model.Other) {
	err = global.DB.Where("id = ?", other.ID).First(&_other).Error
	return err, _other
}

// AutoDelete 判断时间自动执行硬删除
func AutoDelete() {
	//计算删除的数据的时间与目前时间过去了多久
	var other []model.Other
	//查询软删除的数据    查删除的时间
	if err := global.DB.Unscoped().Where("deleted_at").Find(&other); err != nil {
		fmt.Println(err)
	}
	nowTime := time.Now()
	var deleteTime = make([]time.Time, 10)
	for i := 0; i < len(other); i++ {
		deleteTime[i] = other[i].DeletedAt.Time
		//sub()表示时间之间的间隔
		fmt.Printf("已删除%v分钟\n", int(nowTime.Sub(deleteTime[i]).Minutes()))
		//相隔时间超过1秒自动删除
		if int(nowTime.Sub(deleteTime[i]).Seconds()) > 1 {
			if err := global.DB.Unscoped().Where("deleted_at").Delete(&other); err != nil {
				fmt.Println(err)
			}
		}
	}
}
