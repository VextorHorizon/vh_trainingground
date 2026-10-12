package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func checkWebsite(url string) {
	// 1. ตั้ง Timeout หน่อย ไม่ใช่รอเว็บตายไปชาติหน้า
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[❌] %s: เข้าถึงไม่ได้ (Error: %v)\n", url, err)
		return // ทำงานต่อ ไม่ต้องตายตาม
	}
	
	// 2. สำคัญมาก! ต้องปิด Body ทุกครั้ง (ใช้ defer ให้เป็นนิสัย)
	defer resp.Body.Close()

	// 3. เช็คสถานะแบบกว้างๆ (2xx คือ OK)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("[✅] %s is alive! (Status: %d)\n", url, resp.StatusCode)
	} else {
		fmt.Printf("[⚠️] %s might be down (Status: %d)\n", url, resp.StatusCode)
	}
}

func main() {
	// 4. เช็ค Argument ก่อนจะเรียกใช้ index 1
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}

	userInput := os.Args[1]
	checkWebsite(userInput)
}
