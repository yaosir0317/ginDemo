package main

import (
	"context"
	"flag"
	"fmt"
	"ginDemo/config"
	"ginDemo/logger"
	"ginDemo/middlewares"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"ginDemo/controllers"
)

var (
	env        int
	configPath string
)

func parseFlag() {
	flag.IntVar(&env, "env", 0, "env")
	flag.StringVar(&configPath, "config", "config/test.yaml", "configPath")
	flag.Parse()
}

func main() {
	parseFlag()
	cfg := config.NewConfig(configPath)
	if err := logger.InitLogger(cfg.Log); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	r := gin.New()

	loggerZap, _ := zap.NewProduction()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(loggerZap, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(loggerZap, true))
	/*
		gin.Default()默认使用了Logger和Recovery中间件，其中：
		Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
		Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	*/
	r.Use(middlewares.CostMiddleWare())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/query", controllers.HandleQuery)
		apiGroup.POST("/form", controllers.HandleForm)
		apiGroup.POST("/json", controllers.HandleJson)
	}
	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Minute,
		IdleTimeout:       1 * time.Second,
		Handler:           r,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			//log.Error("start server error", "err", err.Error())
			os.Exit(1)
		}
	}()
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		//log.Error("http server shutdown failed", "err", err.Error())
		return
	}
}
