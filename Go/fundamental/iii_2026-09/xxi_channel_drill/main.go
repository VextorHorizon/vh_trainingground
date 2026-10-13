package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	for i := 1; i <= 3; i++ {
		ch <- i * 10
	}

	a := <-ch
	b := <-ch
	c := <-ch

	fmt.Println("Total:", a+b+c)
}
