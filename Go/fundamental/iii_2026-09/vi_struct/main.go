package main

import "fmt"

type Player struct {
	Name string
	Age  int
}

func main() {
	player := Player{Name: "O2", Age: 39}

	fmt.Println(player.Name)
}
