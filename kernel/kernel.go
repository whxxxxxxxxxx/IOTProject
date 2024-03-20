package kernel

import (
	"IOTProject/config"
	"IOTProject/store/mysql"
	"IOTProject/store/rds"
	"context"
	"github.com/gin-gonic/gin"
)

type (
	Engine struct {
		GIN *gin.Engine

		SKLMySQL  *mysql.Orm
		MainCache *rds.Redis

		Ctx    context.Context
		Cancel context.CancelFunc

		ConfigListener []func(*config.GlobalConfig)
	}
)

var Kernel *Engine
