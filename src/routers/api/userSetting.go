package api

import (
	"github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/e"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/util"
	"github.com/gin-gonic/gin"
)

type inputStatus struct {
	BackgroundSetting int `json:"background_setting"`
}

func ModifyUserSetting(c *gin.Context, userId int) (int, interface{}) {
	tempInput := new(inputStatus)
	err := c.BindJSON(tempInput)
	if err != nil || tempInput.BackgroundSetting > 2 || tempInput.BackgroundSetting < 0 {
		return util.MakeErrorReturn(e.BadRequest, 40000, e.GetMsg(40000))
	}
	err = models.ModifyUserSetting(userId,tempInput.BackgroundSetting)
	if err != nil {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}

	return util.MakeSuccessReturn(200, gin.H{
		"BackgroundSetting": tempInput.BackgroundSetting,
	})
}

func GetUserSetting(c *gin.Context, userId int) (int, interface{}) {
	temp, err := models.GetUserSetting(userId)
	if err != nil {
		return util.MakeErrorReturn(e.NotFound, 40440, e.GetMsg(40440))
	}

	return util.MakeSuccessReturn(200, gin.H{
		"BackgroundSetting": temp.BackgroundSetting,
	})
}
