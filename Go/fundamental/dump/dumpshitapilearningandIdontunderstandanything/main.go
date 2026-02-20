package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	r.GET("/api/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "เป้าหมายถูกทำลายเรียบร้อย!",
			"data": gin.H{
				"target_id": id,
				"name":      "VextorHorizon",
				"role":      "Future DevOps",
			},
		})
	})
}
