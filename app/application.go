package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	router   = gin.Default()
	database *gorm.DB
)

func StartApplication(db *gorm.DB) {

	database = db
	MapUrls()
	router.Run(":12003")
}
