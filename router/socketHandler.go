package router

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

func socketHandler(conn *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = conn.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("rx: %s", msg)

		if err = conn.WriteMessage(mt, msg); err != nil {
			log.Print("tx:", err)
			break
		}
	}
}
