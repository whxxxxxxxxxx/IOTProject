package router

import (
	"IOTProject/internal/app/device/handler"
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

	deviceGroup := e.Group("/device")
	{
		deviceGroup.POST("", handler.CreateDevice)              // 创建设备
		deviceGroup.DELETE("/:id", handler.DeleteDevice)        // 删除设备
		deviceGroup.PUT("/:id", handler.UpdateDevice)           // 更新设备
		deviceGroup.GET("/:id", handler.GetDevice)              // 获取单个设备信息
		deviceGroup.GET("", handler.ListDevices)                // 获取所有设备信息
		deviceGroup.POST("/page", handler.ListDevicesPage)      // 分页获取设备信息
		deviceGroup.POST("/search/page", handler.SearchDevices) //分页搜索设备信息
	}

	e.GET("/device/v1/err", func(c *gin.Context) {
		response.HTTPFail(c, 500000, "device Init test error", errors.New("this is err"))
	})
}
