package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var cnt = 0

func GetIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		cnt++
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"visitors": strconv.Itoa(cnt),
		})
	}
}
