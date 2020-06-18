package app

import (
	"github.com/gin-gonic/gin"
	"github.com/samderlust/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

//StartApplication start the app
func StartApplication() {
	mapUrls()
	logger.Info("about to start application")
	router.Run(":8080")
}
