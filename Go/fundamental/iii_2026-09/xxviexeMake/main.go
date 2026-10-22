package main

import (
	"fmt"
	"time"
)

func main() {

	time.Now()

	ch := make(chan int, 2)

	go func() {
		car := 1
		for car = range ch {
			<-ch
			fmt.Printf("car %d in at", car)

		}

		close(ch)
	}()

	ch <- 1
	ch <- 2
	ch <- 3
}
