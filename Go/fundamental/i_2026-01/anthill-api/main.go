package main

import (
	"net/http"
	"strconv" // เอาไว้แปลงข้อความเป็นตัวเลข

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ยันต์กันผี (CORS) อนุญาตให้หน้าเว็บโยนเลขมาให้ได้
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// ด่านรับตัวเลข: ใช้ Query Parameters (?a=10&b=5&op=add)
	r.GET("/api/calc", func(c *gin.Context) {

		// 1. รับค่าที่หน้าเว็บส่งมา (มันจะมาเป็นตัวหนังสือ String)
		aStr := c.Query("a")
		bStr := c.Query("b")
		op := c.Query("op") // ชนิดการคำนวณ: "add" หรือ "sub"

		// 2. แปลงข้อความเป็นตัวเลข (String to Integer)
		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)
		result := 0

		// 3. ✨ THE LOGIC ZONE ✨ (คิดเลขตรงนี้แหละ!)
		if op == "add" {
			result = a + b
		} else if op == "sub" {
			result = a - b
		}

		// 4. พ่นผลลัพธ์กลับไปให้หน้าเว็บ
		c.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"operation": op,
			"result":    result,
		})
	})

	r.Run(":8085")
}
