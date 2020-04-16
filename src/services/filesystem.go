package services

import (
	"errors"
	"os"
	"path/filepath"

	"istorage/config"
	"istorage/logger"
	"istorage/models"
)

func GetFileStoragePath(filename string) string {
	cwd, _ := os.Getwd()

	return filepath.Join(cwd, config.Config.Storage.Path, filename)
}

func AbsolutePath(file *models.MediaFile) string {
	return GetFileStoragePath(file.FileSystemPath())
}

func Check(file *models.MediaFile) error {
	if !fileExists(AbsolutePath(file)) {
		return errors.New("File not found")
	}
	return nil
}

func RemoveFile(file *models.MediaFile) error {
	if err := Check(file); err != nil {
		return err
	}

	path := GetFileStoragePath(file.Path)

	logger.Infof("Path %s was cleaned", path)

	return os.RemoveAll(path)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
