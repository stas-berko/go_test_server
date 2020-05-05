package main

import (
	"github.com/gin-gonic/gin"
	"testTask/cmd/server/handler"
)


func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", handler.GetIndex())

	router.Run(":8080")
}
