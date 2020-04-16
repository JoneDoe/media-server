package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"istorage/attachment"
	"istorage/logger"
	"istorage/models"
	"istorage/services"
	"istorage/utils"
)

func StoreAttachment(c *gin.Context) {
	errMsg := "Undefined any file"

	if c.GetHeader("Content-Length") == "0" {
		utils.Response{c}.Error(http.StatusBadRequest, errMsg)

		return
	}

	form, _ := c.MultipartForm()

	files := form.File["files[]"]
	if len(files) == 0 {
		utils.Response{c}.Error(http.StatusBadRequest, errMsg)

		return
	}

	var filesList = make([]models.OutputModel, 0)

	for _, file := range files {
		attach := models.Create(file)

		fm, err := attachment.FileManagerFactory(attachment.FileManagerConfig{
			attach.OriginalFile.MimeType(),
			"original",
		})

		if err != nil {
			logger.Error(err)
			utils.Response{c}.Error(http.StatusBadRequest, fmt.Sprintf("Upload error: %q", err.Error()))

			return
		}

		// Upload the file to specific dst.
		go store(c, attach, fm)

		filesList = append(filesList, models.OutputModel{
			FileName: attach.OriginalFile.Filename(),
			Uuid:     attach.Uuid,
		})
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "files": filesList})
}

func store(c *gin.Context, a *models.Attachment, fm attachment.FileManager) {
	fm.SetFile(a.OriginalFile)

	a.Path = fm.DirManager().Path
	a.Version = fm.ToJson().FileName

	c.SaveUploadedFile(a.OriginalFile.Upload, fm.Filepath())

	services.InitDb().CreateRecord(a)
}
