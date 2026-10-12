package main

import "fmt"

type Motorcycle struct {
    Model string
    Color string
}

func main() {
    // 1. ออกรถมาเป็นสีดำ (Initialization)
    myR7 := Motorcycle{
        Model: "Yamaha R7",
        Color: "Black",
    }
    fmt.Println("Before:", myR7.Color) // Black

    // 2. เอาไปทำสีใหม่เป็นสีม่วง (Direct Assignment)
    myR7.Color = "Purple" 
    
    fmt.Println("After:", myR7.Color)  // Purple (ทับไปโต้งๆ เลย!)
}
