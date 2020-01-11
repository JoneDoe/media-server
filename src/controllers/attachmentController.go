package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"istorage/attachment"
	"istorage/config"
	"istorage/models"
	"istorage/services"
	"istorage/utils"
)

func StoreAttachment(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files[]"]

	var filesList = make([]models.OutputModel, 0)

	for _, file := range files {
		attach, err := attachment.Create(config.Config.Storage.Path, file)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.OnError(fmt.Sprintf("Upload error: %q", err.Error())))
			return
		}

		fm := attachment.NewFileManager(attachment.FileManagerConfig{
			attach.Dir,
			attach.OriginalFile.BaseMime,
			"original",
		})
		fm.SetFilename(attach.OriginalFile.Ext())

		// Upload the file to specific dst.
		if err = c.SaveUploadedFile(file, fm.Filepath()); err != nil {
			c.JSON(http.StatusBadRequest, utils.OnError(fmt.Sprintf("Upload error: %q", err.Error())))
			return
		}

		filesList = append(filesList, models.OutputModel{
			FileName: attach.OriginalFile.Filename,
			Uuid:     attach.Uuid,
		})

		attach.Version = fm.ToJson().FileName

		go services.InitDb().CreateRecord(attach)
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "files": filesList})
}

// Get parameters for convert from Request query string
func GetConvertParams(req *http.Request) (map[string]string, error) {
	raw_converts := req.URL.Query().Get("converts")

	if raw_converts == "" {
		raw_converts = "{}"
	}

	convert := make(map[string]string)

	err := json.Unmarshal([]byte(raw_converts), &convert)
	if err != nil {
		return nil, err
	}

	return convert, nil
}
