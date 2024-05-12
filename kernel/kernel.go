package kernel

import (
	"IOTProject/config"
	"IOTProject/store/mysql"
	"IOTProject/store/tdengine"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type (
	Engine struct {
		GIN *gin.Engine

		SKLMySQL *mysql.Orm
		TDEngine *tdengine.Orm

		Ctx    context.Context
		Cancel context.CancelFunc

		JSCmd *exec.Cmd

		HttpServer *http.Server

		CurrentIpList []string

		ConfigListener []func(*config.GlobalConfig)
	}
)

var Kernel *Engine
