package main

import (
	"fmt"
)

type Guitar struct {
	Brand string
	Price float32
}

func (g Guitar) Info() {
	fmt.Printf("This is a %s, Price: %.2f\n", g.Brand, g.Price) // ใช้ Printf ใส่ \n ด้วย อย่าลืม
}

func (g *Guitar) Upgrade(newPrice float32) {
	g.Price = newPrice // := คือประกาศตัวแปรใหม่ ส่วน = คือการให้ค่าตัวแปรเดิมที่มีอยู่แล้ว
}

func main() {

	myGuitar := Guitar{Brand: "Music Man", Price: 120000}

	myGuitar.Info()
	myGuitar.Upgrade(134000)

	fmt.Println("----Update Price----")
	myGuitar.Info()
}
