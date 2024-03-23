package router

import (
	"IOTProject/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func AppDeviceInit(e *gin.Engine) {
	e.GET("/device/v1", func(c *gin.Context) {
		response.HTTPSuccess(c, map[string]any{
			"message": "device Init Success",
		})
	})

	e.GET("/device/v1/err", func(c *gin.Context) {
		response.HTTPFail(c, 500000, "device Init test error", errors.New("this is err"))
	})
}
