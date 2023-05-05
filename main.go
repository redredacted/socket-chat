package main

import (
	"github.com/redredacted/socket-chat/router"
)

func main() {
	app := router.NewApplication()
	app.SetupMiddleWare()
	app.SetupRoutes()
	app.Listen(":3000")
}
