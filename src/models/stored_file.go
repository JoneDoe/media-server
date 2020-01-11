package models

type File struct {
	Dir, Name, Type string
	Versions        Version
}

type Version struct {
	Original struct {
		Filename, Url string
		Size          int
	}
}

type FSFile struct {
	FileName string `json:"fileName"`
	Url      string `json:"url"`
}
