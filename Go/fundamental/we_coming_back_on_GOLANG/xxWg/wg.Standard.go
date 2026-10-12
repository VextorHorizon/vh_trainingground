package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("Worker:", i)
		}() //<- same as something(input)<- // it mean func(input type){}(inputFromUserToThisFunction)
	}

	wg.Wait()
	fmt.Println("Done")

}
