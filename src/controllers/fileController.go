package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"istorage/models"
	"istorage/services"
	"istorage/utils"
)

func ReadFile(c *gin.Context) {
	token, _ := bindRequestToken(c)

	resp := utils.Response{c}

	rec, err := services.InitDb().GetRecord(token.Uuid)
	if err != nil {
		resp.Error(http.StatusNotFound, "File not found")
		return
	}

	media := models.InitMedia(rec)

	if err = services.Check(media); err != nil {
		resp.Error(http.StatusNotFound, "Can`t read file")
		return
	}

	c.FileAttachment(services.AbsolutePath(media), media.Name)
}

func DeleteFile(c *gin.Context) {
	token, _ := bindRequestToken(c)

	resp := utils.Response{c}

	rec, err := services.InitDb().GetRecord(token.Uuid)
	if err != nil {
		resp.Error(http.StatusNotFound, "File not found")
		return
	}

	media := models.InitMedia(rec)

	go removeMedia(media, token.Uuid)

	resp.Success(http.StatusOK, token.Uuid)
}

func removeMedia(file *models.MediaFile, uuid string) error {
	if err := services.RemoveFile(file); err != nil {
		return err
	}

	return services.InitDb().DeleteRecord(uuid)
}

func bindRequestToken(c *gin.Context) (*models.RequestToken, error) {
	token := &models.RequestToken{}
	if err := c.ShouldBindUri(&token); err != nil {
		return nil, err
	}

	return token, nil
}
