package handler

import (
	"github.com/gabehamasaki/chat-go/internal/config"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	config *config.Config
}

func NewHandler(logger *zap.Logger, config *config.Config) *Handler {
	return &Handler{logger: logger, config: config}
}
