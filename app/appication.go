package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StarApplication() {
	mapUrl()
	router.Run(":8080")
}
