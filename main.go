package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/click", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "clock",
		})
	})
	r.Run("0.0.0.0:5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
