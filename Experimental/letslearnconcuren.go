package main

import (
	"fmt"
	"time"
)

func saysomething(text string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("[%s] %s: รอบที่ %d\n", time.Now().Format("15:04:05.000"), text, i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	saysomething("Goroutine is working")

	saysomething("I'm a main function")

	fmt.Println("Finished")
}
