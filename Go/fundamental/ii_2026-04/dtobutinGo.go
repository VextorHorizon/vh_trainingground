package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	user := User{
		Name:   "Vextor",
		Age:    1000,
		Gender: "Male",
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(string(data))

	jsonByte, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("Error")
		return
	}

	var user1 User
	err = json.Unmarshal(jsonByte, &user1)
	if err != nil {
		fmt.Println("Couldn't translate:", err)
		return
	}

	fmt.Println(user1)
}
