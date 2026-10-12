package main

import (
	"fmt"
	"os"
)

// 1. นี่คือ 'พิมพ์เขียว' ของข้อมูลโฟลเดอร์ที่เราอยากได้
type FolderInfo struct {
	Name    string
	Path    string
	Size    int64
	ModTime string
}

func main() {
	// 2. สร้าง 'ถังเก็บ' (Slice) ที่จะเก็บร่างแยกของ Struct
	var scanResults []FolderInfo

	// 3. เริ่มสแกน (สมมติสแกนที่โฟลเดอร์ปัจจุบัน ".")
	entries, _ := os.ReadDir(".")

	for _, entry := range entries {
		if entry.IsDir() {
			info, _ := entry.Info()

			// 4. 'ปั๊ม' ร่างแยกออกมาจาก Struct แม่แบบ
			tempFolder := FolderInfo{
				Name:    entry.Name(),
				Path:    "./" + entry.Name(),
				Size:    info.Size(),
				ModTime: info.ModTime().Format("2006-01-02"),
			}

			// 5. โยนร่างแยกลงถังเก็บ
			scanResults = append(scanResults, tempFolder)
		}
	}

	// แสดงผลลัพธ์ทั้งหมดที่สแกนเจอ
	fmt.Printf("สแกนเสร็จแล้ว! เจอทั้งหมด %d โฟลเดอร์\n", len(scanResults))
	for _, f := range scanResults {
		fmt.Printf("- [%s] ขนาด: %d bytes\n", f.Name, f.Size)
	}
}
