package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func request(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode == 200{
		fmt.Println("Website is alive!")
	} else {
		fmt.Println("Website is down?")
	}
	// fmt.Println(response.StatusCode)

}

func main(){

	args := os.Args
	userInput := args[1]

	if userInput == "" {
		fmt.Println("Empty")
	}
	request(userInput)
	
	

}
