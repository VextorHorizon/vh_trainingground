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

	for _, file := range files {
		if !file.IsDir() {
			fmt.Printf("📄 %s\n", file)
		} else {
			fmt.Printf("📂 %s\n", file)
		}
	}

}
