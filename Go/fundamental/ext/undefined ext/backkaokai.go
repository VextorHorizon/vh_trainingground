package main

import (
	"fmt"
)

func main() {

	isFighting := true
	days := 3

	if isFighting {
		fmt.Printf("เก่งมาก ไปสู้ตั้ง %v วันแนะ\n", days)
	} else {
		fmt.Println("ยังไม่เห็นได้สู้อะไรเล่นนิหว่า")
	}
}
