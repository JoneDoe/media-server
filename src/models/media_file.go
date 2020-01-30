package models

import (
	"path/filepath"

	"github.com/mitchellh/mapstructure"
)

type MediaFile struct {
	Path    string `json:"path"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

func InitMedia(dbRecord interface{}) *MediaFile {
	media := &MediaFile{}
	mapstructure.Decode(dbRecord, media)

	return media
}

func (file *MediaFile) FileSystemPath() string {
	return filepath.Join(file.Path, file.Version)
}
