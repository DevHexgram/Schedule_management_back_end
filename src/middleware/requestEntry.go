package middleware

import (
	"github.com/DevHexgram/Schedule_management_back_end/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RequestEntryWithStatus(f func(c *gin.Context, userId uint) (int, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		//tempOwner,exist:=c.Get("owner")
		//strOwner := fmt.Sprintf("%v",tempOwner)
		strId := c.GetString("userId")
		intId, err := strconv.Atoi(strId)
		if err != nil {
			c.JSON(util.MakeErrorReturn(500, 50020, "Middleware Wrong"))
		}

		c.JSON(f(c, uint(intId)))
	}
}

func RequestEntryDefault(f func(c *gin.Context) (int, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(f(c))
	}
}
