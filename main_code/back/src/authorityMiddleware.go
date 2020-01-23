package main

import (
	"github.com/gin-gonic/gin"
)

//needAuthority:所需要的最低权限
func authorityControl(needAuthority int) gin.HandlerFunc {
	return func(c *gin.Context) {
		authority := c.GetInt("authority")
		if authority < needAuthority {
			c.JSON(makeErrorReturn(400,40040,"Poor Authority"))
			c.Abort()
			return
		}
		c.Next()
	}
}
