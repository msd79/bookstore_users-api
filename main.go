package main

import (
	"github.com/msd79/bookstore_users-api/app"
	"github.com/msd79/bookstore_users-api/logger"
)

func main() {
	app.StartApplication()
	logger.Info("Starting application")
}
