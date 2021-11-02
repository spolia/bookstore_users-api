package app

import (
	"github.com/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StarApplication() {
	mapUrl()

	logger.Info("start the application...")
	router.Run(":8080")
}
