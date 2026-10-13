// ทายก่อนรัน: จบยังไง? แล้วถ้าเปลี่ยน send(ch) เป็น go send(ch) ล่ะ?
package main

import "fmt"

func send(ch chan int) {
	ch <- 10
}

func main() {
	ch := make(chan int)
	send(ch)
	fmt.Println(<-ch)
}
