package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Context *gin.Context
}

func OnError(msg string) map[string]interface{} {
	return gin.H{
		"status": "error",
		"error":  msg,
	}
}

func (r Response) ErrorMsg(msg string) {
	r.Context.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error":  msg,
	})
}

func (r Response) SuccessMsg(msg map[string]interface{}) {
	r.Context.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   msg,
	})
}
