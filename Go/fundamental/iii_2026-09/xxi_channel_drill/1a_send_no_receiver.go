// ทายก่อนรัน: พิมพ์อะไรออกมา และจบปกติหรือ deadlock?
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("worker: ก่อนส่ง")
		ch <- 1
		fmt.Println("worker: ส่งเสร็จ")
	}()
	time.Sleep(time.Second)
	fmt.Println("main: จบ")
}
