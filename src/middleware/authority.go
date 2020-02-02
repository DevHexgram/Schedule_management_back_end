package middleware

import (
	"github.com/DevHexgram/Schedule_management_back_end/pkg/util"
	"github.com/gin-gonic/gin"
)

//needAuthority:所需要的最低权限
func AuthorityControl(needAuthority int) gin.HandlerFunc {
	return func(c *gin.Context) {
		authority := c.GetInt("authority")
		if authority < needAuthority {
			c.JSON(util.MakeErrorReturn(400,40040,"Poor Authority"))
			c.Abort()
			return
		}
		c.Next()
	}
}
