package server

import (
	"IOTProject/config"
	"IOTProject/internal/app/appInitialize"
	"IOTProject/kernel"
	"IOTProject/store/mysql"
	"IOTProject/store/rds"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configYml string
	engine    *kernel.Engine
	StartCmd  = &cobra.Command{
		Use:     "server",
		Short:   "Set Application config info",
		Example: "main server -c config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setUp()
			loadStore()
			loadApp()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yaml", "Start server with provided configuration file")
}

// 初始化配置和日志
func setUp() {
	// 初始化全局 ctx
	ctx, cancel := context.WithCancel(context.Background())

	// 初始化资源管理器
	engine = &kernel.Engine{Ctx: ctx, Cancel: cancel}
	kernel.Kernel = engine

	// 加载配置
	config.LoadConfig(configYml, func(globalConfig *config.GlobalConfig) {
		for _, listener := range engine.ConfigListener {
			listener(globalConfig)
		}
	})

	//初始化Gin
	mode := config.GetConfig().MODE
	if mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 初始化 Gin
	engine.GIN = gin.New()
	// 使用内置的 Recovery 中间件，用于恢复任何panic，如果有panic的话
	engine.GIN.Use(gin.Recovery())
	engine.GIN.Use(cors.Default())
	engine.GIN.Use(gin.Logger())

}

// 存储介质连接
func loadStore() {
	engine.SKLMySQL = mysql.MustNewMysqlOrm(config.GetConfig().SKLMysql)
	engine.MainCache = rds.MustNewRedis(config.GetConfig().MainCache)
}

// 加载应用，包含多个生命周期
func loadApp() {
	apps := appInitialize.GetApps()
	for _, app := range apps {
		_err := app.PreInit(engine)
		if _err != nil {
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.Init(engine)
		if _err != nil {
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.PostInit(engine)
		if _err != nil {
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.Load(engine)
		if _err != nil {
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.Start(engine)
		if _err != nil {
			os.Exit(1)
		}
	}

}

// 启动服务
func run() {
	port := config.GetConfig().Port
	// 开启 tcp 监听
	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logx.SystemLogger.Errorw("failed to listen", zap.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
	}

	go func() {
		if _err := engine.Grpc.Serve(grpcL); _err != nil {
			logx.SystemLogger.Errorw("failed to start to listen and serve grpc", zap.Field{Key: "error", Type: zapcore.StringType, String: _err.Error()})
		}
	}()

	go func() {
		logx.SystemLogger.Info("mux listen starting...")
		if _err := tcpMux.Serve(); _err != nil {
			logx.SystemLogger.Errorw("failed to serve mux", zap.Field{Key: "error", Type: zapcore.StringType, String: _err.Error()})
		}
	}()

	println(stringx.Green("Server run at:"))
	println(fmt.Sprintf("-  Local:   http://localhost:%s", port))
	localHost := ip.GetLocalHost()
	engine.CurrentIpList = make([]string, 0, len(localHost))
	for _, host := range localHost {
		engine.CurrentIpList = append(engine.CurrentIpList, host)
		println(fmt.Sprintf("-  Network: http://%s:%s", host, port))
	}
	// 健康检查设置为可接受服务
	healthz.Health.Set(true)

	// 监听退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 健康检查设置为不可接受服务
	healthz.Health.Set(false)

	println(stringx.Blue("Shutting down server..."))
	//tracex.StopAgent()

	if engine.SlsClient != nil {
		if err = engine.SlsClient.Close(); err != nil {
			println(stringx.Yellow("Sls client close failed: " + err.Error()))
		}
	}
	logx.SystemLogger.Stop()
	logx.ServiceLogger.Stop()

	ctx, cancel := context.WithTimeout(engine.Ctx, 5*time.Second)
	defer engine.Cancel()
	defer cancel()

	if err := engine.HttpServer.Shutdown(ctx); err != nil {
		println(stringx.Yellow("Server forced to shutdown: " + err.Error()))
	}

	println(stringx.Green("Server exiting Correctly"))
}
