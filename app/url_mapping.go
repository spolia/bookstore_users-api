package app

import (
	"github.com/bookstore_users-api/controllers/ping"
	"github.com/bookstore_users-api/controllers/user"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", user.GetUser)
	router.GET("/user/search", user.SearchUser)
	router.POST("/users", user.CreateUser)
}
