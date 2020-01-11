package attachment

import (
	"mime/multipart"
	"strings"

	"github.com/google/uuid"

	"istorage/models"
	"istorage/upload"
)

// Attachment contain info about directory, base mime type and all files saved.
type Attachment struct {
	OriginalFile  *upload.OriginalFile
	Dir           *DirManager
	Versions      map[string]FileManager
	Version, Uuid string
}

type FileManagerConfig struct {
	Dir               *DirManager
	MimeType, Version string
}

// Function receive root directory, original file, convert params.
// Return Attachment saved.
func Create(storage string, file *multipart.FileHeader) (*Attachment, error) {
	mime := strings.Split(file.Header.Get("Content-Type"), "/")[0]

	dm, err := CreateDir(storage, mime)
	if err != nil {
		return nil, err
	}

	ofile := &upload.OriginalFile{mime, "", file.Filename, file.Size}

	attachment := &Attachment{
		OriginalFile: ofile,
		Dir:          dm,
		Versions:     make(map[string]FileManager),
		Uuid:         uuid.New().String(),
	}

	return attachment, nil
}

func (a *Attachment) ApplyConverts(converts map[string]string) {
	for version, convert_opt := range converts {
		fm, err := a.CreateVersion(version, convert_opt)
		if err != nil {
			continue
		}

		a.Versions[version] = fm
	}
}

// Directly save single version and return FileManager.
func (attachment *Attachment) CreateVersion(version string, convert string) (FileManager, error) {
	fm := NewFileManager(FileManagerConfig{attachment.Dir, attachment.OriginalFile.BaseMime, version})
	fm.SetFilename(attachment.OriginalFile.Ext())

	if err := fm.Convert(attachment.OriginalFile.Filepath, convert); err != nil {
		return nil, err
	}

	return fm, nil
}

func (attachment *Attachment) ToJson() *models.MediaFile {
	return &models.MediaFile{
		Dir:     attachment.Dir.Path,
		Name:    attachment.OriginalFile.Filename,
		Type:    attachment.OriginalFile.BaseMime,
		Version: attachment.Version,
	}
	/*data := make(map[string]interface{})
	data["type"] = attachment.OriginalFile.BaseMime
	data["dir"] = attachment.Dir.FileSystemPath
	data["name"] = attachment.OriginalFile.Filename

	versions := make(map[string]interface{})
	for version, fm := range attachment.Versions {
		versions[version] = fm.ToJson()
	}
	data["versions"] = versions

	return data*/
}
