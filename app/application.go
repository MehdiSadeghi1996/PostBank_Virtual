package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	router = gin.Default()

	database *gorm.DB
)

func StartApplication(db *gorm.DB) {

	// Disable CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	database = db
	MapUrls()
	router.Run(":12003")
}
