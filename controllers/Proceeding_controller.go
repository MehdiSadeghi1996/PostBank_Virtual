package controllers

import (
	"PostBank_Virtual_Banking/domain"
	"PostBank_Virtual_Banking/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ProceedingController struct {
	proceedingRepository *repository.ProceedingRepository
}

func NewProceedingController(proceedingRepository *repository.ProceedingRepository) *ProceedingController {
	return &ProceedingController{proceedingRepository: proceedingRepository}
}

func (pc *ProceedingController) Create(c *gin.Context) {

	var newProceeding domain.Proceeding
	if err := c.BindJSON(&newProceeding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := pc.proceedingRepository.CreateProceeding(&newProceeding)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &newProceeding)
}

func (pc *ProceedingController) GetById(c *gin.Context) {

	idAsString := c.Param("id")
	id, err := strconv.ParseUint(idAsString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	gama, err := pc.proceedingRepository.GetProceedingByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gama)
}

func (pc *ProceedingController) GetPagination(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))

	gamas, err := pc.proceedingRepository.GetProceedingWithPagination(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gamas)
}

func (pc *ProceedingController) GetFilters(c *gin.Context) {
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

	proceedingList, err := pc.proceedingRepository.MultipleColumnFilter(filters, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filters":        filters,
		"proceedingList": proceedingList,
	})
}
