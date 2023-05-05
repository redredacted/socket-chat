package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/redredacted/socket-chat/middleware"
)

type Application struct {
	inner *fiber.App
}

func NewApplication() *Application {
	return &Application{inner: fiber.New()}
}

func (app *Application) SetupMiddleWare() {
	app.inner.Use("/ws", middleware.WsUpgrade)
}

func (app *Application) SetupRoutes() {
	app.inner.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("rx: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Print("tx:", err)
				break
			}
		}
	}))
}

func (app *Application) Listen(port string) {
	app.inner.Listen(port)
}
