package main

import (
	"encoding/json"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func main() {
	file, _ := os.Create("data.json")
	defer file.Close()

	u := User{ID: 2, Username: "Silvy", Age: 21}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	encoder.Encode(u)
}
