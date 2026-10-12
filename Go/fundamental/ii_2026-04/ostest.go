package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	target, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(string(target))

	var data = map[string]any{}
	json.Unmarshal(target, &data)

	fmt.Println(data)
	namefromJson := data["name"]

	fmt.Println(namefromJson)

}
