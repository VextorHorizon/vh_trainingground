// ทายก่อนรัน: พิมพ์อะไรออกมา และจบปกติหรือ deadlock?
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("worker")
	}()
	ch <- 10
	fmt.Println(<-ch)
}
