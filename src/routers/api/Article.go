package api

import (
	"fmt"
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

type modifyArticle struct {
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

	//tagId, err := models.AddTag(tempInput.TagName, username, uint(userId))
	//if err != nil || tagId <= 0 {
	//	return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	//}
	//
	//ok := models.AddArticle(tempInput.ArticleBody, tagId, uint(userId))
	//if ok == false {
	//	return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	//}
	ok := models.AddTagAndArticle(tempInput.TagName, username, &tempInput.ArticleBody, uint(userId))
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}
	return util.MakeSuccessReturn(200, "")
}

func GetArticle(c *gin.Context, userId int) (int, interface{}) {
	StrTagId := c.Query("tag_id")
	tagId, err := strconv.Atoi(StrTagId)
	if err != nil {
		return util.MakeErrorReturn(e.BadRequest, 40400, e.GetMsg(40400))
	}
	out := models.GetArticle(tagId, userId)

	return util.MakeSuccessReturn(200, out)
}

func ModifyArticle(c *gin.Context, userId int) (int, interface{}) {
	strTagId := c.Query("tag_id")
	tagId, err := strconv.Atoi(strTagId)
	if err != nil {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}

	tempModifyArticle := new(modifyArticle)
	err = c.BindJSON(tempModifyArticle)
	if err != nil {
		return util.MakeErrorReturn(e.BadRequest, 40000, e.GetMsg(40000))
	}
	ok := models.ModifyArticleWithTag(&tempModifyArticle.ArticleBody, tagId, userId)
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}
	return util.MakeSuccessReturn(e.Success, "")
}

func ModifyTag(c *gin.Context, userId int) (int, interface{}) {
	strTagId := c.Query("tag_id")
	tagName := c.Query("tag_name")
	fmt.Println(tagName)
	fmt.Println(strTagId)
	tagId, err := strconv.Atoi(strTagId)

	if err != nil {
		return util.MakeErrorReturn(e.BadRequest, 40400, e.GetMsg(40400))
	}
	ok := models.ModifyTag(tagName, tagId, userId)
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}
	return util.MakeSuccessReturn(e.Success, "")
}

func DeleteArticle(c *gin.Context, userId int) (int, interface{}) {
	strTagId := c.Query("tag_id")
	tagId, err := strconv.Atoi(strTagId)
	if err != nil {
		return util.MakeErrorReturn(e.NotFound, 40400, e.GetMsg(40400))
	}
	isFound := models.FindTag(tagId, userId)
	if isFound == false {
		return util.MakeErrorReturn(e.NotFound, 40430, e.GetMsg(40430))
	}
	ok := models.DeleteTagAndArticle(tagId, userId)
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}
	return util.MakeSuccessReturn(e.Success, "")
}
