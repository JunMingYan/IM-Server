package socketIO

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/models"
)

func RunSocketIO(c *gin.Context) {
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

	Server := socketio.NewServer(nil)

	Server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println(s.ID())
		return nil
	})

	Server.OnEvent("/", "sendValidateMessage", func(s socketio.Conn, ws models.WS) {
		fmt.Println(ws)
		fmt.Println("test")
	})

	go Server.Serve()
	defer Server.Close()

}
