// ทายก่อนรัน: เหมือน 1a แต่ไม่มี Sleep — พิมพ์อะไรออกมา และจบปกติหรือ deadlock?
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("worker: ก่อนส่ง")
		ch <- 1
		fmt.Println("worker: ส่งเสร็จ")
	}()
	fmt.Println("main: จบ")
}
