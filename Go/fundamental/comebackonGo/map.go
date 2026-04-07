package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"age":  17,
		"year": 2008,
	}

	fmt.Println(m["age"])
}
