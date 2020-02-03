package models

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

type imageURL struct {
	gorm.Model
	URL string
}

//获取随机图片
func GetRandomImage() string {
	var sum int
	DB.Table("image_urls").Count(&sum)

	rand.Seed(time.Now().Unix())
	tempID := rand.Intn(sum-1)+1

	tempImageURL := new(imageURL)
	DB.Table("image_urls").Where("id = ?",tempID).Find(tempImageURL)

	return tempImageURL.URL
}