package utils

import (
	"github.com/gin-gonic/gin"
)

func OnError(msg string) map[string]interface{} {
	return gin.H{
		"status": "error",
		"error":  msg,
	}
}
