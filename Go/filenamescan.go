package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err) // ถ้า error ให้ log ว่า error อะไร แล้วก้กระโดดออกหน้าต่าง
	}
	// fmt.Printf("%v", files)
	// fmt.Println(len(files))
	var foldbox []string
	var docbox []string
	for _, file := range files {
		if file.IsDir() {
			foldbox = append(foldbox, file.Name())
		} else {
			docbox = append(docbox, file.Name())
		}
	}

	for _, file := range foldbox {
		fmt.Printf("📂 - %s\n", file)
	}
	for _, file := range docbox {
		fmt.Printf("📄 - %s\n", file)
	}
}
