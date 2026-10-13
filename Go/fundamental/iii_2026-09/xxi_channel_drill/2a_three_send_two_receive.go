// ทายก่อนรัน: พิมพ์อะไรออกมา และจบปกติหรือ deadlock? ตอน end ค่าตัวที่ 3 อยู่ที่ไหน?
package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 1; i <= 3; i++ {
		go func() { ch <- i }()
	}
	a := <-ch
	b := <-ch
	fmt.Println(a + b)
	fmt.Println("end")
}
