package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Password  string
	Authority int
}

//注册
func AddNewUser(password string, username string, authority int) (tempUser *User, ok bool) {
	tx := DB.Begin()
	tempUser = &User{
		//Model:     gorm.Model{},
		Username:  username,
		Password:  password,
		Authority: authority,
	}
	if tx.Create(tempUser).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return nil, ok
	}
	tx.Commit()
	ok = true
	return tempUser, ok
}

//登陆,返回 ok==false 代表失败
func FindUser(password string, username string) (tempUser *User, ok bool) {
	DB.Where("username = ? AND password = ?", username, password).Find(tempUser)
	if tempUser.ID <= 0 {
		ok = false
		return
	}
	ok = true
	return
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
