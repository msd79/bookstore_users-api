package app

import (
	"github.com/gin-gonic/gin"
	"github.com/msd79/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

//StartApplication ...
func StartApplication() {
	mapUrls()

	logger.Info("Starting application")
	router.Run(":8080")

}
