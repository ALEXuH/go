package input

import "errors"

type FileInput struct {
}

func (fi *FileInput) retrieveData() ([]byte, error) {
	return nil, errors.New("")
}

func NewFileInput() *FileInput {
	return &FileInput{}
}
