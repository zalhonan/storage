package main

import (
	"fmt"

	"github.com/zalhonan/storage/internal/storage"
)

func main() {
	fmt.Println("ok, it works, what now?")

	st := storage.NewStorage()

	fmt.Println(st)

}
