package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	port := os.Getenv("PORT")
	if len(port) == 0 { 
		port = "3000" 
	}

	mapUrls()
	router.Run(":" + port)
}
