package models

import (
	"mime/multipart"

	"github.com/google/uuid"
)

// Attachment contain info about original uploaded file, uuid...
type Attachment struct {
	OriginalFile        *OriginalFile
	Path, Version, Uuid string
}

// Return Attachment
func Create(file *multipart.FileHeader) *Attachment {
	return &Attachment{
		OriginalFile: &OriginalFile{file},
		Uuid:         uuid.New().String(),
	}
}

func (attachment *Attachment) ToJson() *MediaFile {
	return &MediaFile{
		Path:    attachment.Path,
		Name:    attachment.OriginalFile.Filename(),
		Type:    attachment.OriginalFile.MimeType(),
		Version: attachment.Version,
	}
}
