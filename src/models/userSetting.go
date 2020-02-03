package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type userSetting struct {
	gorm.Model
	UserID           uint
	BackgroundSetting int //0:color ; 1:URL image ; 2:customize image
}

//获取用户设置
func GetUserSetting(userId int) (*userSetting,error) {
	temp := new(userSetting)
	DB.Table("user_statuses").Where("user_id = ?", userId).Find(temp)
	if temp.ID <= 0 {
		return nil,errors.New("NotFound")
	}
	return temp,nil
}

func ModifyUserSetting(userId int,backgroundSetting int) error{
	tx := DB.Begin()
	if tx.Model(&userSetting{}).Where("user_id = ?", userId).Update(&userSetting{
		BackgroundSetting: backgroundSetting,
	}).RowsAffected != 1 {
		tx.Rollback()
		return errors.New("Can'tInsertIntoDatabase")
	}
	tx.Commit()
	return nil
}
