package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"strconv"
	"testTask/cmd/server/utils"
)

func GetIndex(jsonStorage *os.File) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := utils.GetFromStorage(jsonStorage)
		cnt := resp.Visitors + 1

		newData := utils.NewVisitorsCountData(cnt)
		newData.RewriteStorageData(jsonStorage)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"visitors": strconv.Itoa(cnt),
		})
	}
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	utils.VisitorsDataNotifier.Register(&utils.EventObserver{Conn: conn})

}

func OpenWS() gin.HandlerFunc {
	return func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	}

}
