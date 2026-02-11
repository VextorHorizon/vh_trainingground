package service

import (
	"fmt"
)

func ProcessData(name string) string {
	if name == ""{
		name = "Stranger"
	}
	return fmt.Sprintf("Hello, %s! ข้อมูลถูกประมวลผลผ่าน Service แล้วนะ (๑˃ᴗ˂)ﻭ", name)
}