package api

import (
	"github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/e"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetAllDailyAffairs(c *gin.Context, userId int) (int, interface{}) {
	data := models.GetAllDailyAffair(userId)
	return util.MakeSuccessReturn(e.Success, data)
}

func GetAllAffairs(c *gin.Context, userId int) (int, interface{}) {
	out := models.GetAllAffairs(userId)
	//fmt.Println(out)
	return util.MakeSuccessReturn(e.Success, out)
}
