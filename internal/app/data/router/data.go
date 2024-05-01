package router

import (
	"IOTProject/internal/app/data/handler"
	"IOTProject/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func AppDataInit(e *gin.Engine) {
	e.GET("/data/v1", func(c *gin.Context) {
		response.HTTPSuccess(c, map[string]any{
			"message": "data Init Success",
		})
	})

	e.GET("/data/v1/err", func(c *gin.Context) {
		response.HTTPFail(c, 500000, "data Init test error", errors.New("this is err"))
	})

	dataGroup := e.Group("/data")
	{
		dataGroup.GET("/:id", handler.GetOneData) // 获取单个数据信息
	}
}
