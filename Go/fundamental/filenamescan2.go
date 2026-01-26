package main

import (
	"fmt"
	"log"
	"os"
)

func listFiles(dirPath string) {
	// ใช้ os.ReadDir เพื่ออ่าน Directory ข้อมูลที่ได้จะเป็น []os.DirEntry
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("เห้ย! เข้าไปอ่านโฟลเดอร์ไม่ได้: %v", err)
	}

	fmt.Printf("--- กำลังเช็คไฟล์ใน: %s ---\n", dirPath)

	for _, entry := range entries {
		// เช็คว่าเป็น Directory หรือเปล่า
		if entry.IsDir() {
			fmt.Printf("[DIR]  %s\n", entry.Name())
		} else {
			// ถ้าเป็นไฟล์ปกติ
			info, _ := entry.Info() // ดึงข้อมูลเพิ่มเติม (Size, ModTime)
			infoInMB := float32(info.Size()) / (1024 * 1024)
			if info.Size() < (1024 * 1024) {
				fmt.Printf("[FILE] %-20s | Size: %d bytes\n", entry.Name(), info.Size())
			} else {
				fmt.Printf("[FILE] %-20s | Size: %g Megabytes\n", entry.Name(), infoInMB)
			}
		}
	}
}

func main() {
	// ลองใส่ path ที่ต้องการเช็คดู
	targetDir := "."
	listFiles(targetDir)
}
