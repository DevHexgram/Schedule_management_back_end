package models

import "github.com/jinzhu/gorm"

type InvitationCode struct {
	gorm.Model
	Code string
}

type imageURL struct {
	gorm.Model
	URL string
}