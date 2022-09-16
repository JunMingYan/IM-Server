package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/models"
)

var up = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunWebSocket(c *gin.Context) {
	var user models.UserWS
	err := c.Bind(&user)
	if err != nil {
		logrus.Errorf("解析参数出错,err:%s", err)
		c.JSON(http.StatusOK, gin.H{
			"msg":  "无效参数，请重新输入",
			"code": "400",
		})
		return
	}

	ws, err := up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		User: user,
		Conn: ws,
		Send: make(chan []byte, 1024),
	}

	MyServer.Register <- client
	go client.Read()
	go client.Write()
}
