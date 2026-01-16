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
		// fmt.Println(info)

		filemb := (float64(info.Size()) / 1024 / 1024)
		if filemb < 1 {
			fmt.Printf("File Name: %-25s File Size: %10d Byte\n", info.Name(), info.Size())
		} else {
			fmt.Printf("File Name: %-25s File Size: %10.2f Megabyte\n", info.Name(), filemb)
		}
	}

}
