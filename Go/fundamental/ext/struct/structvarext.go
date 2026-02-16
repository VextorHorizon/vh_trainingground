package main

import (
	"fmt"
)

type UserConfig struct {
	Username string
	Level    int
	isAdmin  bool
}

func main() {
	var FirstUser UserConfig

	FirstUser.Username = "VextorHorizon"
	FirstUser.Level = 23
	FirstUser.isAdmin = true

	secondUser := UserConfig{
		Username: "Silvy",
		Level:    34,
		isAdmin:  false,
	}

	fmt.Println(FirstUser.Username, "vs", secondUser.Username)
}
