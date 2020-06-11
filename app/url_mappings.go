package app

import (
	"github.com/samderlust/bookstore_users-api/controllers/ping"
	"github.com/samderlust/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.POST("/users/search", users.SeachUser)
	router.GET("/users/:id", users.GetUser)
}
