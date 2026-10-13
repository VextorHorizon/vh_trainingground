// ทายก่อนรัน: พิมพ์อะไรออกมา และจบปกติหรือ deadlock?
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("worker: ทำงาน")
	}()
	fmt.Println(<-ch)
	fmt.Println("main: จบ")
}
