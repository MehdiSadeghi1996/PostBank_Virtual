package controllers

import (
	"PostBank_Virtual_Banking/domain"
	"PostBank_Virtual_Banking/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type GamaController struct {
	gamaRepository *repository.GamaRepository
}

func NewGamaController(gamaRepository *repository.GamaRepository) *GamaController {
	return &GamaController{gamaRepository: gamaRepository}
}

func (gc *GamaController) Create(c *gin.Context) {

	var newGama domain.Gama
	if err := c.BindJSON(&newGama); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := gc.gamaRepository.Create(&newGama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &newGama)
}

func (gc *GamaController) GetById(c *gin.Context) {

	idAsString := c.Param("id")
	id, err := strconv.ParseUint(idAsString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	gama, err := gc.gamaRepository.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gama)
}

func (gc *GamaController) GetPagination(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))

	gamas, err := gc.gamaRepository.GetWithPagination(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gamas)
}

func (gc *GamaController) GetFilters(c *gin.Context) {
	filters := make(map[string]interface{})

	for key, values := range c.Request.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
		return
	}

	// Set default end time 24 hours later if end time is not provided
	var endTime time.Time
	if endTimeStr != "" {
		endTime, err = time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format"})
			return
		}
	} else {
		endTime = startTime.Add(24 * time.Hour)
	}

	gamaList, err := gc.gamaRepository.MultipleColumnFilter(filters, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filters":  filters,
		"gamaList": gamaList,
	})
}
