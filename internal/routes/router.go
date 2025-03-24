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
	api := engine.Group("/api")
	{
		api.GET("/chats", r.handler.ListChats)
		api.GET("/healthy", r.handler.Healthy)
	}
}
