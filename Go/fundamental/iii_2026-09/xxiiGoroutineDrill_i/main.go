package main

import (
	"fmt"
	"sync"
	"time"
)

func slowDouble(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	ch <- n * 2

}

func main() {
	var wg sync.WaitGroup

	timeStart := time.Now()
	ch := make(chan int)

	nums := []int{1, 2, 3, 4}
	var numsResult []int

	for _, num := range nums {
		wg.Add(1)                   // สร้างตามจำนวน value ที่อยู่ใน slice nums
		go slowDouble(num, ch, &wg) // เปิด columm แล้วไปต่อเลย | ในกรณีนี้เปิด 4 ตัวเพราะว่า len(nums) = 4 (มี value 4 ตัวนั้นหล่ะเอาง่ายๆ)
	} // range ให้ครบจำนวน value ใน slice nums

	for i := 1; i <= len(nums); i++ { // ในขณะเดียวกันหลังจากเปิด columm ทั้ง 4 เสร็จแล้ว มาที่บรรทัดนี้ต่อเลย
		numsResult = append(numsResult, <-ch) // เปิดรอ? ถูกไหมนะ <- ch เปิด channel รอจนกว่าจะมี ตัวมาจับคู่ส่งของ | ส่งของเสร็จ append เข้า slice numsResult และไปบรรทัดต่อไปซึ่งคือ loop
	}
	wg.Wait() // รอให้ทั้ง 4 ตัวทำงานเสร็จ ก้คือ ส่งค่า, บรรทัด numsResult รอจับคู่, จับคู่ได้, append, ไปบรรทัดต่อไป, loop ทำงานเสร็จ, ไม่มี worker เหลือ, เรียบร้อยกลับบ้าน

	fmt.Println(numsResult)
	fmt.Println(time.Since(timeStart))
}
