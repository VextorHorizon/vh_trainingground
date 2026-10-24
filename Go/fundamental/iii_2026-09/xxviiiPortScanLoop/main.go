package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {

	for i := 1; i <= 100; i++ {

		portString := strconv.Itoa(i)
		host := "192.168.1.1:" + portString

		result, err := net.Dial("tcp", host)

		if err != nil {
			fmt.Printf("Port is closed: %s\n", host)
		}
		if result != nil {
			result.Close()
			fmt.Printf("Port is open: %s\n", host)
		}
	}

}
