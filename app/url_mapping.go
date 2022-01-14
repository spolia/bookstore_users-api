package app

import (
	"github.com/bookstore_users-api/controllers/ping"
	"github.com/bookstore_users-api/controllers/user"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", user.Get)
	router.POST("/users", user.Create)
	router.POST("/users/login", user.Login)
	router.PUT("/user/:user_id", user.Update)
	router.PATCH("/user/:user_id", user.Update)
	router.DELETE("/user/:user_id", user.Delete)
}
