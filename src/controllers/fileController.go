package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"istorage/imaginary"
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

	resizeProfile, _ := bindResizeProfile(c)
	if resizeProfile.ProfileName != "" {
		resizeProfile.MediaFile = media
		c.Set("resizeProfile", resizeProfile)
		return
	}

	c.FileAttachment(services.AbsolutePath(media), media.Name)
}

func ReadFileWithResize(c *gin.Context) {
	resizeProfile := c.MustGet("resizeProfile").(*imaginary.ResizeProfile)

	pattern := strings.Join([]string{"cropper", ".*", filepath.Ext(resizeProfile.MediaFile.Name)}, "")
	tmpFile, _ := ioutil.TempFile("", pattern)
	defer os.Remove(tmpFile.Name()) // clean up

	err := imaginary.Resize(resizeProfile.ProfileName, services.AbsolutePath(resizeProfile.MediaFile), tmpFile.Name())
	if err != nil {
		utils.Response{c}.Error(http.StatusNotFound, strings.Join([]string{
			"Can`t make operation, try one of following: ",
			imaginary.AvailableProfiles(),
		}, ""))
		return
	}

	c.FileAttachment(tmpFile.Name(), "resized"+filepath.Ext(resizeProfile.MediaFile.Name))
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

func bindResizeProfile(c *gin.Context) (*imaginary.ResizeProfile, error) {
	data := &imaginary.ResizeProfile{}
	if err := c.ShouldBindUri(&data); err != nil {
		return nil, err
	}

	return data, nil
}
