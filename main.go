package main

import (
	"github.com/redredacted/socket-chat/router"
	"os"
)

func main() {
	app := router.NewApplication()
	app.SetupMiddleWare()
	app.SetupRoutes()
	app.Listen(os.Getenv("PORT"))
}
