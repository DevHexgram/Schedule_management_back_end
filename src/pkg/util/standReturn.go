package util

import "github.com/gin-gonic/gin"

func MakeSuccessReturn(status int, data interface{}) (int, interface{}) {
	return status, gin.H{
		"error": 0,
		"msg":   "success",
		"data":  data,
	}
}

func MakeErrorReturn(status int, error int, msg string) (int, interface{}) {
	return status, gin.H{
		"error": error,
		"msg":   msg,
	}
}

