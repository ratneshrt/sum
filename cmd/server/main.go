package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/sum", func(c *gin.Context) {
		a, err := strconv.Atoi(c.Query("a"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid or missing parameter a",
			})
		}

		b, err := strconv.Atoi(c.Query("b"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid or missing parameter a",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"result": a + b,
		})
	})

	r.Run(":3000")
}
