package main

import (
	"github.com/gin-gonic/gin"
	"testTask/cmd/server/handler"
	"testTask/cmd/server/utils"
)





func main() {
	router := gin.Default()
	jsonStorage := utils.InitStorage("storage/visitors.json")

	defer jsonStorage.Close()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", handler.GetIndex(jsonStorage))
	router.GET("/ws", handler.OpenWS())
	err := router.Run(":8080"); utils.Check(err)
}
