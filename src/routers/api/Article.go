package api

import (
	"github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/e"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"unicode/utf8"
)

type inputArticle struct {
	TagName     string `json:"tag_name"`
	ArticleBody string `json:"article_body"`
}

func GetAllTag(c *gin.Context, userId int) (int, interface{}) {
	out := models.GetAllTags(userId)
	return util.MakeSuccessReturn(e.Success, out)
}

func AddArticle(c *gin.Context, userId int) (int, interface{}) {
	username := c.GetString("username")
	tempInput := new(inputArticle)

	err := c.BindJSON(tempInput)
	if err != nil || tempInput.TagName == "" || utf8.RuneCountInString(tempInput.ArticleBody) > 1500 {
		return util.MakeErrorReturn(e.BadRequest, 40000, e.GetMsg(40000))
	}

	tagId, err := models.AddTag(tempInput.TagName, username, uint(userId))
	if err != nil || tagId <= 0 {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}

	ok := models.AddArticle(tempInput.ArticleBody, tagId, uint(userId))
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}

	return util.MakeSuccessReturn(200, "")
}

func GetArticle(c *gin.Context, userId int) (int, interface{}) {
	StrTagId := c.Query("tagId")
	tagId, err := strconv.Atoi(StrTagId)
	if err != nil {
		return util.MakeErrorReturn(e.BadRequest, 40400, e.GetMsg(40400))
	}
	out := models.GetArticle(tagId, userId)

	return util.MakeSuccessReturn(200, out)
}

//func ModifyArticle(c *gin.Context, userId int) (int, interface{}) {
//
//}
//

//func DeleteArticle(c *gin.Context, userId int) (int, interface{}) {
//	c.Query("id")
//}
