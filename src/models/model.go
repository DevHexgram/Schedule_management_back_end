package models

import (
	"fmt"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	strDb := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Config.DB.User,
		setting.Config.DB.Password,
		setting.Config.DB.Address,
		setting.Config.DB.DBName)

	DB, err = gorm.Open("mysql", strDb)

	if err != nil {
		panic(err)
	}

	fmt.Println("success connect to database")

	DB.AutoMigrate(&affair{})
	DB.AutoMigrate(&dailyAffair{})

	DB.AutoMigrate(&InvitationCode{})
	DB.AutoMigrate(&imageURL{})

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&userSetting{})

	DB.AutoMigrate(&ArticleTag{})
	DB.AutoMigrate(&Article{})
}
