package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	ch <- 1
	fmt.Println(13)
	fmt.Println(<-ch)
}
