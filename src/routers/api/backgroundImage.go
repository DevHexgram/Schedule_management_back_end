package api

import (
	"github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context) {
	URL := models.GetRandomImage()
	c.Redirect(e.Found, URL)
}
