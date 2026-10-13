// ทายก่อนรัน: พิมพ์อะไรออกมา และจบปกติหรือ deadlock?
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 10
	time.Sleep(100 * time.Millisecond)
}
