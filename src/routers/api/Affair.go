package api

import (
	"errors"
	"github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/e"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type input struct {
	Title    string    `json:"title" binding:"required"`
	Deadline time.Time `json:"deadline"`
	Extra    string    `json:"extra"`
}

func AddAffair(c *gin.Context, userId int) (int, interface{}) {
	temp := new(input)
	err := c.BindJSON(temp)
	if err != nil {
		return util.MakeErrorReturn(e.BadRequest, 40000, e.GetMsg(40000))
	}
	ok := models.AddAffair(userId, temp.Title, temp.Extra, temp.Deadline)
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}

	return util.MakeSuccessReturn(200, "")
}

func DeleteAffair(c *gin.Context, userId int) (int, interface{}) {
	strId := c.Query("id")
	if strId == "" {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}

	err = models.DeleteAffair(userId, id)
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

func ModifyAffair(c *gin.Context, userId int) (int, interface{}) {
	strId := c.Query("id")
	if strId == "" {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}
	ok := models.FindAffair(userId, id)
	if ok == false {
		return util.MakeErrorReturn(e.NotFound, 40430, e.GetMsg(40430))
	}

	temp := new(input)
	err = c.BindJSON(temp)
	if err != nil {
		//fmt.Println(err)
		return util.MakeErrorReturn(e.BadRequest, 40000, e.GetMsg(40000))
	}

	ok = models.ModifyAffair(userId,id,temp.Title,temp.Extra,temp.Deadline)
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}

	return util.MakeSuccessReturn(200, "")
}
