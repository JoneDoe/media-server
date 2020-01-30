package models

type Image struct {
	Dir, Name, Type string
	Versions        struct {
		Original  ImageStruct
		Thumbnail ImageStruct
	}
}

type ImageStruct struct {
	Filename, Url       string
	Size, Height, Width int
}
