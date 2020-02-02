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
	ok := models.AddNewUser(tempUser.Password, tempUser.Username, 3)
	if ok == false {
		return util.MakeErrorReturn(e.InternalServerError,50000,e.GetMsg(50000))
	}

	tempInvitationCode := new(InvitationCode)
	s.DB.Table("invitation_codes").Where("code = ?", tempUser.Code).Find(tempInvitationCode)
	if tempInvitationCode.ID <= 0 {
		return util.MakeErrorReturn(404, 40420, "Invitation Code Wrong")
	}

	dbUser := new(User)
	s.DB.Table("users").Where("username = ?", tempUser.Username).Find(dbUser)
	if dbUser.ID > 0 {
		return util.MakeErrorReturn(400, 40030, "Duplicate username")
	}

	tx := s.DB.Begin()
	if tx.Create(&User{
		Username:  tempUser.Username,
		Password:  tempUser.Password,
		Authority: 3,
	}).RowsAffected != 1 {
		tx.Rollback()
		return util.MakeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()

	s.DB.Table("users").Where("username = ? AND password = ?", tempUser.Username, tempUser.Password).Find(dbUser)

	tx = s.DB.Begin()
	if tx.Create(&userStatus{
		UserID:           dbUser.ID,
		BackgroundStatus: 0,
	}).RowsAffected != 1 {
		tx.Rollback()
		return util.MakeErrorReturn(500, 50000, "Can't Insert Into Database")
	}
	tx.Commit()

	token, err := util.GenerateToken(dbUser.Username, dbUser.Authority, dbUser.ID)

	if err != nil {
		//fmt.Println(err)
		return util.MakeErrorReturn(500, 50010, "Can't Generate Token")
	}
	return util.MakeSuccessReturn(200, gin.H{"token": token})
}

func Login(c *gin.Context) (int, interface{}) {
	tempUser := new(InputUser)
	err := c.BindJSON(tempUser)
	if err != nil || tempUser.Username == "" || tempUser.Password == "" {
		return util.MakeErrorReturn(400, 40000, "Wrong Format of JSON")
	}

	dbUser := new(User)
	s.DB.Table("users").Where("username = ? AND password = ?", tempUser.Username, tempUser.Password).Find(dbUser)
	if dbUser.ID <= 0 {
		return util.MakeErrorReturn(404, 40410, "Username or Password Wrong")
	}

	token, err := util.GenerateToken(dbUser.Username, dbUser.Authority, dbUser.ID)
	if err != nil {
		return util.MakeErrorReturn(500, 50010, "Can't Generate Token")
	}

	return util.MakeSuccessReturn(200, gin.H{"token": token})
}
