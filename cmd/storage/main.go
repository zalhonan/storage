package main

import (
	"fmt"
	"log"

	"github.com/zalhonan/storage/internal/storage"
)

func main() {
	fmt.Println("ok, it works, what now?")

	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("ooh my"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file uploaded", file)

	fileById, err := st.GetById(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("got file by ID", fileById)
}
