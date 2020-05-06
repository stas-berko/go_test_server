package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"testTask/cmd/server/utils"
)

func GetIndex(jsonStorage *os.File) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := utils.GetFromStorage(jsonStorage)
		cnt := resp.Visitors+1

		newData := &utils.VisitorsData{Visitors: cnt}
		newData.RewriteStorageData(jsonStorage)


		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"visitors": strconv.Itoa(cnt),
		})
	}
}
