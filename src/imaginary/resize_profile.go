package imaginary

import "istorage/models"

type ResizeProfile struct {
	ProfileName string `uri:"profile" binding:"required"`
	MediaFile   *models.MediaFile
}
