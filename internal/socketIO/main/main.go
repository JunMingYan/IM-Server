package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
)

func main() {
	r := gin.Default()
	Server := socketio.NewServer(nil)

	Server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println(s.ID())
		return nil
	})

	//Server.OnEvent("/", "sendValidateMessage", func(s socketio.Conn, ws models.WS) {
	//	fmt.Println(ws)
	//	fmt.Println("test")
	//})

	go Server.Serve()
	defer Server.Close()
	//
	r.GET("/IO/socket.io", gin.WrapH(Server))
	r.Run(":9999")
}
