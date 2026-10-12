package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TargetWebsite struct {
	name      string
	url       string
	isitAlive bool
}

func (T *TargetWebsite) Request(targetUrl string) { // Not finished, going to apply struct to this script

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	response, err := client.Get(T.url)
	if err != nil {
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode <= 300 {
		fmt.Printf("✅ Website is online! url: %s \n", T.url)
	} else {
		fmt.Printf("❌ Website is down! url: %s \n", T.url)
	}

}

func RequestAndProcess(url string, wg *sync.WaitGroup) {

	defer wg.Done()

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	responses, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error: %s %s\n", err, url)
		return // อันนี้เขา(ai)บอกมีโอกาสเจ๊ง
	}

	defer responses.Body.Close()

	if responses.StatusCode >= 200 && responses.StatusCode <= 300 {
		fmt.Printf("✅ Website is online! url: %s \n", url)
	} else {
		fmt.Printf("❌ Website is down! url: %s \n", url)
	}

}

func main() {

	var wg sync.WaitGroup
	// content, _ := os.ReadFile("urls.txt")
	urls := []string{
		"https://go.dev/tour/moretypes/12",
		"https://www.print3dd.com/fdm-material-guide/",
		"https://github.com/VextorHorizon/vh_trainingground",
		"https://github.com/VextorHorizon/vh_traininggroundswq",
		"https://www.google.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go RequestAndProcess(url, &wg)
	}

	fmt.Printf("Checking %v website..\n", len(urls))
	wg.Wait()

	fmt.Println("Finish")
}

// List url เว็บ (ซัก 100 เว็บ)
// Request ดึงข้อมูล จาก url เว็บ //err check คือ ถ้าไม่มีเว็บๆนั้นหรือไม่เจอไม่มีข้อมูลจาก Url ให้ print ว่าไม่เจอเว็บนั้น แล้ว skip
// Process เช็คว่าเว็บแต่ละเว็บ Online(status >= 200 && status < 300) หรือป่าว
// ให้แบ่งแยก Print ให้เว็บที่ Online(เขียว) ขึ้นก่อน Offline(แดง)

// list url
// request
// process check if online

// In future I'm going to make this to API
