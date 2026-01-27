package main

import (
	"fmt"
	"net/http"
	"time"
)

func request(url string) {
	response, err := http.Get(url)
}