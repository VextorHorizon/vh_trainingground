package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go printHello(ch)
	message := <-ch

	fmt.Println(message)
}

func printHello(che chan string) {
	che <- "Hello"
}

// <- = Put in pipe
// make(chan type) = make pipe for data transit
// message := <-ch == put value from pipe to message varable
