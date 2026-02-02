package main

import (
	"fmt"
	"net/http"
	"time"
)

type TargetWebsite struct {
	url       string
	isitAlive bool
}

func RequestAndProcess(t *TargetWebsite) (bool, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	response, err := client.Get(t.url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode <= 300 {
		t.isitAlive = true
		return true, err
	} else {
		t.isitAlive = false
		return false, err
	}

}

func main() {

	targetWebsite := TargetWebsite{ // Future: Adding it to be loop slice for next time it will scan txt file and pack it to struct one by one
		url:       "https://www.google.com",
		isitAlive: false,
	}

	isAlive, err := RequestAndProcess(&targetWebsite)

	if err != nil { // Future: Is this great?
		fmt.Printf("Website is error with: %s", err)
	}

	if isAlive {
		fmt.Printf("✅ Website is online! url: %s \n", targetWebsite.url)
	} else {
		fmt.Printf("❌ Website is offline! url: %s \n", targetWebsite.url)
	}

}
