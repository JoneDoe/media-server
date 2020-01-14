package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"istorage/attachment"
	"istorage/models"
	"istorage/services"
	"istorage/utils"
)

func StoreAttachment(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files[]"]

	var filesList = make([]models.OutputModel, 0)

	for _, file := range files {
		attach := models.Create(file)

		fm, err := attachment.FileManagerFactory(attachment.FileManagerConfig{
			attach.OriginalFile.MimeType(),
			"original",
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, utils.OnError(fmt.Sprintf("Upload error: %q", err.Error())))
			return
		}

		fm.SetFile(attach.OriginalFile)

		attach.Path = fm.DirManager().Path
		attach.Version = fm.ToJson().FileName

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
	c.SaveUploadedFile(a.OriginalFile.Upload, fm.Filepath())
	services.InitDb().CreateRecord(a)
}
