package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"istorage/models"
	"istorage/services"
	"istorage/utils"
)

func ReadFile(c *gin.Context) {
	file, err := initRequestFile(c)
	if err != nil {
		utils.Response{c}.ErrorMsg(err.Error())
		return
	}

	rec, err := services.InitDb().GetRecord(file.Uuid)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	media := models.InitMedia(rec)

	if err = services.Check(media); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.FileAttachment(services.AbsolutePath(media), media.Name)
}

func DeleteFile(c *gin.Context) {
	file, err := initRequestFile(c)
	if err != nil {
		utils.Response{c}.ErrorMsg(err.Error())
		return
	}

	rec, err := services.InitDb().GetRecord(file.Uuid)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	media := models.InitMedia(rec)

	if err = services.Check(media); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	utils.Response{c}.SuccessMsg(rec)
}

func initRequestFile(c *gin.Context) (*models.ReqFile, error) {
	file := &models.ReqFile{}
	if err := c.ShouldBindUri(&file); err != nil {
		return nil, err
	}

	return file, nil
}
