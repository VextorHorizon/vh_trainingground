// ทายก่อนรัน: worker กี่ตัวพิมพ์ "ส่งแล้ว"? ตัวที่ไม่พิมพ์ถือค่าอะไรอยู่?
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 1; i <= 3; i++ {
		go func() {
			ch <- i
			fmt.Println("worker", i, "ส่งแล้ว")
		}()
	}
	a := <-ch
	b := <-ch
	time.Sleep(100 * time.Millisecond)
	fmt.Println("a =", a, "b =", b)
	fmt.Println("end")
}
