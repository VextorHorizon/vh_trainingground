package main

import (
	"fmt"
	"net/http"
	"my-app/service" // Import service module ของเราเอง
)

func main() {
	// 1. กำหนด Route: เมื่อมีคนเรียกมาที่ "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// รับ Input จาก URL Query (เช่น /?name=Silvy)
		input := r.URL.Query().Get("name")

		// 2. เรียกใช้ Service (นี่ไงที่แกอยากได้!)
		result := service.ProcessData(input)

		// 3. พิมพ์ผลลัพธ์ออกทาง Browser/Client
		fmt.Fprint(w, result)
		
		// พิมพ์ Log ใน Terminal ให้เราเห็นด้วย
		fmt.Printf("มีคนส่ง input: '%s' เข้ามานะจ๊ะ\n", input)
	})

	// 4. สั่งรัน Service บน Port 8080
	port := ":8080"
	fmt.Printf("Server เริ่มทำงานแล้วที่ http://localhost%s ... อย่าทำพังล่ะ! (￢_￢)\n", port)
	
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}