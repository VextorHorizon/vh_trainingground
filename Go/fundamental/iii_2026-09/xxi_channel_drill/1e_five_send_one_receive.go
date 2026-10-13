// ทายก่อนรัน: ตอน end พิมพ์ออกมา มี worker ส่งไม่สำเร็จกี่ตัว และจบปกติหรือ deadlock?
package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 1; i <= 5; i++ {
		go func() { ch <- i }()
	}
	fmt.Println(<-ch)
	fmt.Println("end")
}
