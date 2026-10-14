package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func slowUpper(s string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	ch <- strings.ToUpper(s)
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan string)

	timeStart := time.Now()

	words := []string{"go", "rust", "zig", "c", "odin", "ruby", "typescript"}
	var resultWords []string

	for _, word := range words {
		wg.Add(1)
		go slowUpper(word, &wg, ch)
	}

	go func() { wg.Wait(); close(ch) }()

	for v := range ch { // v <- ch
		resultWords = append(resultWords, v)
	}

	fmt.Println(resultWords)
	fmt.Println(time.Since(timeStart))

}

/*
ตอนแรกชั้นเอา ไปวางไว้ใต้ go func() { wg.Wait(); close(ch) }()
แล้วมันก้ deadlock มาไล่ดู อ่อ range ch รอตัวต่อไปอยู่มันทำงานแบบ sequential มันไปไม่ถึง close(ch) เราเลยขยับขึ้นมาข้างบน range
ส่วนสาเหตุว่าทำไม wg.Wait(); close(ch) ต้องไปแยกเป็น goroutine แยก
เพราะว่า ตัวมันเองทำงานแบบ sequential และเจ้า <- ch หรือจะ ch<- มันก้รอทำงานทั้งนั้น
กรณีแรก ถ้าเอา wg.Wait(); close(ch) ไปไว้เหนือ range ch{} เราก้จะได้ว่า ch <- strings.ToUpper(s) รอส่งอยู่แต่ไปไม่ถึงเพราะว่า wg.Wait() อยู่
กรณีสองถ้าเอาไปอยู่ใต้ range ก้จะได้ว่า range ch ทำงานจน value ใน ch หมด แต่ว่ามันก้ยังทำงานต่อไปอยู่และไม่มีค่าเข้ามาเพิ่ม มันก้รออยู่อย่างงั้นและ deadlock
*/
