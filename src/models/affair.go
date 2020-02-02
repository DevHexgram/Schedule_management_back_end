package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type affair struct {
	gorm.Model
	Title    string
	Deadline time.Time
	Extra    string
	//Owner string
	UserId uint
}

type dailyAffair struct {
	gorm.Model
	Title string
	Extra string
	//Owner string
	UserId uint
}

