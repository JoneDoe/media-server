package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"bitbucket.org/vadimtitov/istorage/attachment"
	"bitbucket.org/vadimtitov/istorage/config"
	"bitbucket.org/vadimtitov/istorage/upload"
	"bitbucket.org/vadimtitov/istorage/utils"
)

func StoreAttachment(c *gin.Context) {
	converts, err := GetConvertParams(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.OnError(fmt.Sprintf("Query params: %s", err)))
		return
	}

	converts["original"] = ""

	pavo, _ := c.Request.Cookie("istore")
	if pavo == nil {
		pavo = &http.Cookie{
			Name:    "istore",
			Value:   uuid.New().String(),
			Expires: time.Now().Add(10 * 356 * 24 * time.Hour),
			Path:    "/",
		}
		c.Request.AddCookie(pavo)
		http.SetCookie(c.Writer, pavo)
	}

	iStorage := config.Config.Storage.Path

	files, err := upload.Process(c.Request, iStorage)
	if err == upload.Incomplete {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"file":   gin.H{"size": files[0].Size},
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.OnError(fmt.Sprintf("Upload error: %q", err.Error())))
		return
	}

	data := make([]map[string]interface{}, 0)
	for _, ofile := range files {
		attachment, err := attachment.Create(iStorage, ofile, converts)
		if err != nil {
			data = append(data, map[string]interface{}{
				"name":  ofile.Filename,
				"size":  ofile.Size,
				"error": err.Error(),
			})
			continue
		}
		data = append(data, attachment.ToJson())
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "files": data})
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
