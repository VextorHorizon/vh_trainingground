package main

import (
	"flag"
	"fmt"
)

func main() {

	p := flag.String("p", "", "help message")

	flag.Parse()

	fmt.Printf("Hello!: %s", *p)

}
