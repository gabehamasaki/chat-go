package routes

import (
	"github.com/gabehamasaki/chat-go/internal/config"
	"github.com/gabehamasaki/chat-go/internal/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	logger  *zap.Logger
	config  *config.Config
	handler *handler.Handler
}

func NewRouter(logger *zap.Logger, handler *handler.Handler, config *config.Config) *Router {
	return &Router{logger: logger, handler: handler, config: config}
}

func (r *Router) Setup(engine *gin.Engine) {
	r.logger.Info("Config", zap.Any("config", r.config))
	api := engine.Group("/api")
	{
		api.GET("/healthy", r.handler.Healthy)
	}
}
