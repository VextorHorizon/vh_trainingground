package main

import (
	"encoding/json"
	"fmt"
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

	jsonStr := `{"name":"Vextor","age":1000,"gender":"Male"}`

	var user1 User
	json.Unmarshal([]byte(jsonStr), &user1)

	fmt.Println(user1)
}
