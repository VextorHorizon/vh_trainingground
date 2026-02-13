package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		return
	}
	fmt.Println(files)
}
