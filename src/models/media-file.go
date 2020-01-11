package models

import (
	"path"

	"github.com/mitchellh/mapstructure"
)

type MediaFile struct {
	Dir     string `json:"dir"`
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
	return path.Join(file.Dir, file.Version)
}
