package main

import (
	"github.com/gin-gonic/gin"
)

type inputStatus struct {
	BackgroundStatus int `json:"background_status"`
}

func (s *Service) modifyUserStatus(c *gin.Context, userId uint) (int, interface{}) {
	tempInput := new(inputStatus)
	err := c.BindJSON(tempInput)
	if err != nil || tempInput.BackgroundStatus > 2 || tempInput.BackgroundStatus < 0 {
		return makeErrorReturn(400, 40000, "Wrong Format Of Json")
	}

	tx := s.DB.Begin()
	if tx.Model(&userStatus{}).Where("user_id = ?", userId).Update(&userStatus{
		BackgroundStatus: tempInput.BackgroundStatus,
	}).RowsAffected != 1 {
		tx.Rollback()
		return makeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()
	return makeSuccessReturn(200,gin.H{
		"BackgroundStatus":tempInput.BackgroundStatus,
	})
}

func (s *Service) getUserStatus(c *gin.Context, userId uint) (int, interface{}) {
	temp := new(userStatus)
	s.DB.Table("user_statuses").Where("user_id = ?", userId).Find(temp)
	if temp.ID <= 0 {
		return makeErrorReturn(404, 40440, "User Status Not Exist")
	}
	return makeSuccessReturn(200, gin.H{
		"BackgroundStatus": temp.BackgroundStatus,
	})
}
