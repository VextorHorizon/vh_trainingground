package main

import (
	"fmt"
)

func main() {
	parking := make(chan string, 2)

	parking <- "1"
	fmt.Println("คันที่ 1 เข้าแล้ว")
	parking <- "2"
	fmt.Println("คันที่ 2 เข้าแล้ว")
	<-parking
	fmt.Println("คัน 1 หรือ 2 ออกแล้ว")
	parking <- "3"
	fmt.Println("คันที่ 3 เข้าแล้ว")
}
