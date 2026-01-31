package main

import "fmt"

// 1. นี่คือ 'พิมพ์เขียว' (Struct) ของเรา
// นิยามว่า 'รถหนึ่งคัน' ในระบบของเราต้องมีข้อมูลอะไรบ้าง
type Motorcycle struct {
	Model string
	Color string
	Km    int
}

// 2. นี่คือ 'Method' (พฤติกรรมที่ผูกกับ Struct)
// ใช้ Pointer (*) เพื่อให้สามารถแก้ไขข้อมูลในกล่องจริงได้ (ไม่ใช่แค่ก๊อปปี้)
func (m *Motorcycle) Paint(newColor string) {
	fmt.Printf("--- กำลังเปลี่ยนสีรถจาก %s เป็น %s ---\n", m.Color, newColor)
	m.Color = newColor
}

// 3. นี่คือ Function ทั่วไปที่ 'รับค่า' เป็น Struct
// เอาไว้ใช้จัดการ Logic ภายนอกที่ไม่ใช่พฤติกรรมหลักของวัตถุ
func CheckServiceStatus(m Motorcycle) {
	if m.Km > 10000 {
		fmt.Printf("[!] แจ้งเตือน: %s วิ่งมา %d km แล้ว ต้องเข้าศูนย์ด่วน!\n", m.Model, m.Km)
	} else {
		fmt.Printf("[✓] สถานะ: %s ยังวิ่งได้กริบๆ อยู่\n", m.Model)
	}
}

func main() {
	// --- การใช้งานจริง ---

	// สร้าง Object (ร่างแยก) จาก Struct
	myR7 := Motorcycle{
		Model: "Yamaha R7",
		Color: "Black",
		Km:    12000,
	}

	// เรียกดูข้อมูลเบื้องต้น
	fmt.Printf("รถของฉัน: %s สี %s\n", myR7.Model, myR7.Color)

	// ลองส่งเข้า Function เพื่อเช็คระยะ
	CheckServiceStatus(myR7)

	// ลองใช้ Method เพื่อเปลี่ยนสีรถ (เป็นสีม่วงที่แกชอบไง)
	myR7.Paint("Purple")

	// เช็คผลลัพธ์หลังเปลี่ยนสี
	fmt.Println("สีใหม่ของรถ:", myR7.Color) // Output: Purple

	// --- แถม: ถ้ามีรถหลายคัน (Slice) แบบที่สงสัย ---
	garage := []Motorcycle{
		myR7,
		{Model: "Ducati Panigale", Color: "Red", Km: 2000},
	}

	fmt.Println("\n--- รายชื่อรถในโรงรถ (Garage) ---")
	for i, bike := range garage {
		fmt.Printf("%d. %s [%s]\n", i+1, bike.Model, bike.Color)
	}
}
