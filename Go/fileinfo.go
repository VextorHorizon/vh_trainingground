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
		info, err := file.Info()
		if err != nil {
			continue
		}
		fmt.Println(info)
		fmt.Println("File Size: ", info.Size(), "Byte")
		filemb := (float64(info.Size()) / 1024 / 1024)
		fmt.Printf("File Size in mb %.2f Megabyte\n", filemb)

	}

}
