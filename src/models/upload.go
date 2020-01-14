package models

import (
	"mime/multipart"
	"path/filepath"
	"strings"
)

// Structure describes the state of the original file.
type OriginalFile struct {
	Upload *multipart.FileHeader
}

func (origin *OriginalFile) MimeType() string {
	return strings.Split(origin.Upload.Header.Get("Content-Type"), "/")[0]
}

func (origin *OriginalFile) Filename() string {
	return origin.Upload.Filename
}

func (origin *OriginalFile) Ext() string {
	return strings.ToLower(filepath.Ext(origin.Filename()))
}
