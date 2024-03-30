package app

import (
	"PostBank_Virtual_Banking/controllers"
	"PostBank_Virtual_Banking/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapUrls() {

	proceedingRepository := repository.NewProceedingRepository(database)
	proceedingController := controllers.NewProceedingController(proceedingRepository)

	gamaRepository := repository.NewGamaRepository(database)
	gamaController := controllers.NewGamaController(gamaRepository)

	router.POST("/api/v1/proceeding", proceedingController.Create)
	router.POST("/api/v1/gama", gamaController.Create)

	router.GET("/api/v1/proceeding/:id", proceedingController.GetById)
	router.GET("/api/v1/proceeding", proceedingController.GetPagination)
	router.GET("/api/v1/proceeding/filter", proceedingController.GetFilters)

	router.GET("/api/v1/gama/:id", gamaController.GetById)
	router.GET("/api/v1/gama", gamaController.GetPagination)
	router.GET("/api/v1/gama/filter", gamaController.GetFilters)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
