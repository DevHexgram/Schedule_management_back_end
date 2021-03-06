package api

import (
	"errors"
	"github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/e"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type dailyInput struct {
	Title string `json:"title" binding:"required"`
	Extra string `json:"extra"`
}

func AddDailyAffair(c *gin.Context, userId int) (int, interface{}) {
	tempAffair := new(dailyInput)
	err := c.BindJSON(tempAffair)
	if err != nil || tempAffair.Title == "" {
		return util.MakeErrorReturn(e.BadRequest, 40000, e.GetMsg(40000))
	}

	err = models.AddDailyAffair(userId, tempAffair.Title, tempAffair.Extra)
	if err != nil {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}

	return util.MakeSuccessReturn(200, "")
}

func DeleteDailyAffair(c *gin.Context, userId int) (int, interface{}) {
	strId := c.Query("id")
	if strId == "" {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}
	err = models.DeleteDailyAffair(userId, id)

	if err != nil {
		if err == errors.New("NotFound") {
			return util.MakeErrorReturn(e.NotFound, 40430, e.GetMsg(40430))
		} else if err == errors.New("Can'tInsertIntoDatabase") {
			return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
		} else {
			return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
		}
	}

	return util.MakeSuccessReturn(200, "")
}

func ModifyDailyAffair(c *gin.Context, userId int) (int, interface{}) {
	strId := c.Query("strId")
	if strId == "" {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}
	isFound := models.FindDailyAffair(userId, id)
	if isFound == false {
		return util.MakeErrorReturn(e.NotFound, 40430, e.GetMsg(40430))
	}

	temp := new(dailyInput)
	err = c.BindJSON(temp)
	//fmt.Println(temp)
	if err != nil {
		//fmt.Println(err)
		return util.MakeErrorReturn(e.BadRequest, 40000, e.GetMsg(40000)) //
	}
	ok := models.ModifyDailyAffair(userId, id, temp.Title, temp.Extra)
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}

	return util.MakeSuccessReturn(200, "")
}
