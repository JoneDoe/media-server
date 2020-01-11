package upload

import (
	"errors"
	"path/filepath"
	"strings"
)

// Error Incomplete returned by uploader when loaded non-last chunk.
var Incomplete = errors.New("Incomplete")

// Structure describes the state of the original file.
type OriginalFile struct {
	BaseMime string
	Filepath string
	Filename string
	Size     int64
}

func (ofile *OriginalFile) Ext() string {
	return strings.ToLower(filepath.Ext(ofile.Filename))
}
