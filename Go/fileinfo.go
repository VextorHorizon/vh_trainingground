package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	files, err := os.ReadDir(".") // มันจะให้ค่าออกมาเป็น Object เช่น filenamescan.go ไม่ใช่แค่ชื่อแต่จะเป็นข้อมูลที่อยู่ใน Object นั้นแบบจะได้ ชื่อ วันที่แก้ไขไฟล์ ขนาด แพ็คอยู่ใน Object นั้นๆเลย
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

		filemb := (float64(info.Size()) / 1024 / 1024) // (1024 / 1024) คือ หารเลข byte ด้วย 1024 * 1024 (1,048,576) เพื่อให้ได้เลข Mebibytes (MiB) ซึ่งเป็น Binary หรือว่าเลขฐานสอง 
		if filemb < 1 {
			fmt.Printf("File Name: %-25s File Size: %10d Byte\n", info.Name(), info.Size())
		} else {
			fmt.Printf("File Name: %-25s File Size: %10.2f Megabyte\n", info.Name(), filemb)
		}
	}

}

