package storage

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s Storage) String() string {
	return "I am the Storage"
}
