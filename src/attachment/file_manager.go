package attachment

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"istorage/models"
	"istorage/upload"
)

type FileManager interface {
	Convert(string, string) error
	SetFilename(string) *FileBaseManager
	Store(*upload.OriginalFile) error
	ToJson() models.FSFile
	Filepath() string
}

type FileBaseManager struct {
	Dir      *DirManager
	Version  string
	Filename string
}

// Return FileManager for given base mime and version.
func NewFileManager(cfg FileManagerConfig) FileManager {
	fbm := &FileBaseManager{Dir: cfg.Dir, Version: cfg.Version}
	switch cfg.MimeType {
	case "image":
		return &FileImageManager{FileBaseManager: fbm}
	default:
		return &FileDefaultManager{FileBaseManager: fbm}
	}

	return nil
}

func (fbm *FileBaseManager) Store(file *upload.OriginalFile) error {
	dest, err := os.Create(fbm.Filepath())
	if err != nil {
		return errors.New("failed to create file")
	}

	defer dest.Close()

	src, err := os.Open(file.Filepath)
	if err != nil {
		return errors.New("failed to read file")
	}

	defer src.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return errors.New("failed to copy file")
	}

	os.Remove(file.Filepath)

	return nil
}

func (fbm *FileBaseManager) SetFilename(ext string) *FileBaseManager {
	salt := strconv.FormatInt(seconds(), 36)
	fbm.Filename = fbm.Version + "-" + salt + ext

	return fbm
}

func (fbm *FileBaseManager) Filepath() string {
	return filepath.Join(fbm.Dir.Abs(), fbm.Filename)
}

func (fbm *FileBaseManager) Url() string {
	return filepath.Join(fbm.Dir.Path, fbm.Filename)
}

func (fdm *FileBaseManager) ToJson() models.FSFile {
	return models.FSFile{fdm.Filename, fdm.Url()}
}

func seconds() int64 {
	t := time.Now()
	return int64(t.Hour()*3600 + t.Minute()*60 + t.Second())
}
