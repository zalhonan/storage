package file

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (File, error) {
	fileId, err := uuid.NewUUID()

	if err != nil {
		return File{}, err
	}

	return File{
		ID:   fileId,
		Name: filename,
		Data: blob,
	}, nil
}
