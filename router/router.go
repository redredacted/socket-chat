package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/redredacted/socket-chat/middleware"
)

type Application struct {
	router *fiber.App
}

func NewApplication() *Application {
	return &Application{router: fiber.New()}
}

func (app *Application) Setup() {
	app.setupMiddleware()
	app.setupRoutes()
}

func (app *Application) setupMiddleware() {
	app.router.Use("/ws", middleware.WsUpgrade)
}

func (app *Application) setupRoutes() {
	app.router.Get("/ws/:id", websocket.New(socketHandler))
	app.router.Get("/health", healthCheck)
}

func (app *Application) Listen(port string) {
	app.router.Listen("0.0.0.0:" + port)
}
