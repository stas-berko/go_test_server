package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var cnt = 0

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		cnt++
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"visitors": strconv.Itoa(cnt),
		})
	})

	router.Run(":8080")
}
