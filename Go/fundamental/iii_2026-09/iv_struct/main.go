package main

import (
	"fmt"
)

func main() {

	type Bigbike struct {
		Name     string
		CC       int
		MaxSpeed int
	}

	Yamaha_R1 := Bigbike{Name: "Yamaha_R1", CC: 998, MaxSpeed: 299}

	fmt.Printf("This motocycle is named:%s. CC is %d. Max speed is %dkm/hr", Yamaha_R1.Name, Yamaha_R1.CC, Yamaha_R1.MaxSpeed)
}
