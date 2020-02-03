package models

import "github.com/jinzhu/gorm"

type InvitationCode struct {
	gorm.Model
	Code string
}

func FindInvitationCode(code string) (isFound bool) {
	tempInvitationCode := new(InvitationCode)
	DB.Table("invitation_codes").Where("code = ?", code).Find(tempInvitationCode)
	if tempInvitationCode.ID <= 0 {
		isFound = false
		return
	}
	isFound = true
	return
}
