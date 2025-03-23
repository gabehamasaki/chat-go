package routes

import (
	"github.com/gabehamasaki/chat-go/internal/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	logger  *zap.Logger
	handler *handler.Handler
}

func NewRouter(logger *zap.Logger, handler *handler.Handler) *Router {
	return &Router{logger: logger, handler: handler}
}

func (r *Router) Setup(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.GET("/healthy", r.handler.Healthy)
	}
}
