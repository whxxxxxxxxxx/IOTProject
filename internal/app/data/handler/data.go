package handler

import (
	"IOTProject/internal/app/data/model"
	"github.com/gin-gonic/gin"
)

func GetOneData(c *gin.Context) {
	id := c.Param("id")
	var status []model.Status
	var performanceMetrics []model.PerformanceMetrics

}
