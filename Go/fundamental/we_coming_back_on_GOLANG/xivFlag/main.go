package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("p", 0, "target")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Missing IP")
		return
	}

	ip := args[0]

	fmt.Println("IP:", ip)
	fmt.Println("Port:", *port)
}
