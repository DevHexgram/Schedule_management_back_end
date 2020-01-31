//连接数据库
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string
	Authority int
}

type affair struct {
	gorm.Model
	Title    string
	Deadline time.Time
	Extra    string
	Owner string
	UserId uint
}

type dailyEvent struct {
	gorm.Model
	Title string
	Extra string
	//Owner string
	userId uint
}

type imageURL struct {
	gorm.Model
	URL string
}

type userStatus struct {
	gorm.Model
	UserID uint
	BackgroundStatus int //1:color ; 2:URL image ; 3:customize image
}

type InvitationCode struct {
	gorm.Model
	Code string
}

func (s *Service) DBInit()  {
	strDb := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.Config.DB.User,
		s.Config.DB.Password,
		s.Config.DB.Address,
		s.Config.DB.DBName)
	db,err := gorm.Open("mysql", strDb)
	DealError(err)
	fmt.Println("success connect to database")

	db.AutoMigrate(&affair{})
	db.AutoMigrate(&dailyEvent{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&InvitationCode{})
	db.AutoMigrate(&imageURL{})
	s.DB = db
	//fmt.Println(s.DB)
}
