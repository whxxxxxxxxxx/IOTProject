package device

import (
	"github.com/gin-gonic/gin"
)

func AppProjectInit(e *gin.Engine) {
	e.GET("/device", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "device",
		})
	})
}
