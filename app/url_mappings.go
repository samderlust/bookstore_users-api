package app

import (
	"github.com/samderlust/bookstore_users-api/controllers/ping"
	"github.com/samderlust/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("internal/users/search", users.Search)
	router.GET("/users/:id", users.GetUser)
	router.PUT("/users/:id", users.UpdateUser)
	router.PATCH("/users/:id", users.UpdateUser)
	router.DELETE("/users/:id", users.DeleteUser)
}
