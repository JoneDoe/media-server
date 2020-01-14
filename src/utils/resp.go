package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Context *gin.Context
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (r Response) Error(code int, msg string) {
	r.Context.AbortWithStatusJSON(code, ErrorResponse{"error", msg})
}

func (r Response) Success(code int, msg string) {
	r.Context.JSON(code, SuccessResponse{"ok", msg})
}
