package attachment

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"istorage/config"
	"istorage/models"
)

type FileManager interface {
	DirManager() *DirManager
	SetFile(*models.OriginalFile) *FileBaseManager
	ToJson() models.FSFile
	Filepath() string
}

type FileBaseManager struct {
	Dir      *DirManager
	Version  string
	Filename string
}

type FileManagerConfig struct {
	MimeType, Version string
}

// Return FileManager for given base mime and version.
func FileManagerFactory(cfg FileManagerConfig) (FileManager, error) {
	dm, err := CreateDir(config.Config.Storage.Path, cfg.MimeType)
	if err != nil {
		return nil, err
	}

	fbm := &FileBaseManager{Dir: dm, Version: cfg.Version}

	switch cfg.MimeType {
	case "image":
		return &FileImageManager{FileBaseManager: fbm}, nil
	default:
		return &FileDefaultManager{FileBaseManager: fbm}, nil
	}
}

func (fbm *FileBaseManager) SetFile(file *models.OriginalFile) *FileBaseManager {
	salt := strconv.FormatInt(seconds(), 36)

	fbm.Filename = strings.Join([]string{
		fbm.Version,
		"-",
		salt,
		file.Ext(),
	}, "")

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

func (fdm *FileBaseManager) DirManager() *DirManager {
	return fdm.Dir
}

func seconds() int64 {
	t := time.Now()
	return int64(t.Hour()*3600 + t.Minute()*60 + t.Second())
}
