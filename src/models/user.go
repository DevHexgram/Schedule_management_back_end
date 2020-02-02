package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Password  string
	Authority int
}

type userSetting struct {
	gorm.Model
	UserID           uint
	BackgroundStatus int //0:color ; 1:URL image ; 2:customize image
}

//注册
func AddNewUser(password string, username string, authority int) (ok bool) {
	tx := DB.Begin()
	if tx.Create(&User{
		//Model:     gorm.Model{},
		Username:  username,
		Password:  password,
		Authority: authority,
	}).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	tx.Commit()
	ok = true
	return
}

//登陆
func FindUser(password string, username string) {
	tempUser := new(User)
	DB.Where("username = ? AND password = ?",username,password).Find(tempUser)
	if tempUser.ID <= 0 {

	}
}

//查重,重复了返回true,无重复返回false
func FindRepeatUser(username string) (isRepeat bool) {
	tempUser := new(User)
	DB.Where("username = ?", username).Find(tempUser)
	if tempUser.ID <= 0 {
		isRepeat = false
		return
	}
	isRepeat = true
	return
}
