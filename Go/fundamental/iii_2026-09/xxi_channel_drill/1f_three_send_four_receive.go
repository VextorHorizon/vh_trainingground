// ทายก่อนรัน: พิมพ์อะไรออกมา และจบปกติหรือ deadlock? ใครค้าง?
package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 1; i <= 3; i++ {
		go func() { ch <- i }()
	}
	a := <-ch
	b := <-ch
	c := <-ch
	d := <-ch
	fmt.Println(a, b, c, d)
}
