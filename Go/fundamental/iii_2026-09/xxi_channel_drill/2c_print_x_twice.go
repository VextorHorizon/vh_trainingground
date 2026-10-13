// ทายก่อนรัน: พิมพ์อะไรออกมา? ตอน end มี worker กี่ตัวที่ยังถือค่าไว้ ไม่ได้ส่งใคร?
package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 1; i <= 4; i++ {
		go func() { ch <- i * 10 }()
	}
	x := <-ch
	fmt.Println(x)
	fmt.Println(x)
	fmt.Println("end")
}
