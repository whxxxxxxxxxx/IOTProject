package handler

import (
	"IOTProject/internal/app/data/dao"
	"IOTProject/internal/app/data/model"
	"IOTProject/middleware/response"
	"github.com/gin-gonic/gin"
)

func GetOneStatusData(c *gin.Context) {
	id := c.Param("id")
	var status []model.Status
	statusId := "s" + id
	tmp, err := dao.Data.GetDataByIdAndCurrentTime(statusId, "1h")
	status = tmp.([]model.Status)
	if err != nil {
		response.ServiceErr(c, err)
	}
	response.HTTPSuccess(c, status)
}

func GetOnePerformanceMetricsData(c *gin.Context) {
	id := c.Param("id")
	var performanceMetrics []model.PerformanceMetrics
	performanceMetricsId := "p" + id
	tmp, err := dao.Data.GetDataByIdAndCurrentTime(performanceMetricsId, "1h")
	performanceMetrics = tmp.([]model.PerformanceMetrics)
	if err != nil {
		response.ServiceErr(c, err)
	}
	response.HTTPSuccess(c, performanceMetrics)
}
