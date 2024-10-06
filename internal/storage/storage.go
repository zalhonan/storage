package storage

import "github.com/zalhonan/storage/internal/file"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s Storage) String() string {
	return "I am the Storage"
}

func (s *Storage) Upload(filename string, blob []byte) (file.File, error) {
	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return file.File{}, err
	}

	return newFile, nil
}
