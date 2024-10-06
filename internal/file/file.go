package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &File{
		ID:   fileId,
		Name: filename,
		Data: blob,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File %v with ID %v", f.Name, f.ID)
}
