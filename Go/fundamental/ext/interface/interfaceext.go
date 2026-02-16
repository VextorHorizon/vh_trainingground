package main

import "fmt"

// 1. สัญญา: "ใครจะเป็น Scanner ต้องมีปุ่ม Scan"
type Scanner interface {
	Scan()
}

// 2. ของจริง (Structs)
type Nmap struct{}

func (n Nmap) Scan() { fmt.Println("Nmap ทำงาน!") }

type ZAP struct{}

func (z ZAP) Scan() { fmt.Println("ZAP ทำงาน!") }

func main() {
	// ประกาศตัวแปร Interface "ตัวเดียว" (เหมือนถังเปล่า 1 ใบ)
	var myTool Scanner

	// จังหวะที่ 1: เอา Nmap ใส่ลงไป
	myTool = Nmap{}
	myTool.Scan() // ผลลัพธ์: Nmap ทำงาน!

	// จังหวะที่ 2: เอา ZAP ใส่ลงไป (ในตัวแปรเดิมเป๊ะ!)
	myTool = ZAP{}
	myTool.Scan() // ผลลัพธ์: ZAP ทำงาน!
}
