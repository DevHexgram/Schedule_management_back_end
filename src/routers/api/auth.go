package api

import (
	"github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/e"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/util"
	"github.com/gin-gonic/gin"
)

type InputUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Code      string `json:"code"`
	Authority int    `json:"authority"`
}

func Register(c *gin.Context) (int, interface{}) {
	tempUser := new(InputUser)
	err := c.BindJSON(tempUser)
	if err != nil || tempUser.Username == "" || tempUser.Password == "" {
		return util.MakeErrorReturn(400, 40000, "Wrong Format of JSON")
	}

	//验证邀请码

	//用户名是否重复
	isRepeat := models.FindRepeatUser(tempUser.Username)
	if isRepeat {
		return util.MakeErrorReturn(e.BadRequest, 40030, e.GetMsg(40030))
	}

	//插入数据库
	dBUser,ok := models.AddNewUser(tempUser.Password, tempUser.Username, 3)
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError, 50000, e.GetMsg(50000))
	}

	//生成token
	token,err := util.GenerateToken(dBUser.Username, dBUser.Authority, dBUser.ID)
	if err != nil {
		return util.MakeErrorReturn(e.InternalServerError,50010,e.GetMsg(50010))
	}

	return util.MakeSuccessReturn(200, gin.H{"token": token})
}

func Login(c *gin.Context) (int, interface{}) {
	//解析并验证json
	tempUserMsg := new(InputUser)
	err := c.BindJSON(tempUserMsg)
	if err != nil || tempUserMsg.Username == "" || tempUserMsg.Password == "" {
		return util.MakeErrorReturn(e.BadRequest, 40000, e.GetMsg(40000))
	}

	//验证登陆信息
	dBUser, ok := models.FindUser(tempUserMsg.Password, tempUserMsg.Username)
	if ok == false {
		util.MakeErrorReturn(e.NotFound,40410,e.GetMsg(40410))
	}

	//生成token
	token,err := util.GenerateToken(dBUser.Username, dBUser.Authority, dBUser.ID)
	if err != nil {
		return util.MakeErrorReturn(e.InternalServerError,50010,e.GetMsg(50010))
	}

	return util.MakeSuccessReturn(200, gin.H{"token": token})
}
