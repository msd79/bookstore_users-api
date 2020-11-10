package app

import (
	"github.com/msd79/bookstore_users-api/controllers/ping"
	controllers "github.com/msd79/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
}
