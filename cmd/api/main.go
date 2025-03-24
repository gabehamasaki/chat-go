package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gabehamasaki/chat-go/internal/config"
	"github.com/gabehamasaki/chat-go/internal/handler"
	"github.com/gabehamasaki/chat-go/internal/logger"
	"github.com/gabehamasaki/chat-go/internal/routes"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func server(lc fx.Lifecycle, logger *zap.Logger, router *routes.Router) *gin.Engine {
	gin.SetMode("release")
	r := gin.New()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	logger.Info("Setup router...")
	router.Setup(r)

	srv := &http.Server{Addr: ":8080", Handler: r}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				logger.Error("Failed to start http server", zap.Error(err))
				return err
			}

			go srv.Serve(ln)
			logger.Info("Http server is started", zap.String("address", srv.Addr))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx)
			logger.Info("Http server is stopped", zap.String("address", srv.Addr))
			return nil
		},
	})

	return r
}

func main() {
	app := fx.New(
		fx.Provide(logger.NewLogger, config.NewConfig, handler.NewHandler, routes.NewRouter, server),
		fx.Invoke(func(*gin.Engine) {}),
	)

	app.Run()
}
