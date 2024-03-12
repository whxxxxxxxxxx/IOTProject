package kernel

import "github.com/gin-gonic/gin"

type (
	Engine struct {
		GIN *gin.Engine
	}
)

var Kernel *Engine
