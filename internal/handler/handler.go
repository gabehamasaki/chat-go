package handler

import (
	"github.com/gabehamasaki/chat-go/internal/config"
	"github.com/gabehamasaki/chat-go/internal/database"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	config *config.Config
	db     *database.Queries
}

func NewHandler(logger *zap.Logger, config *config.Config, db *database.Queries) *Handler {
	return &Handler{logger: logger, config: config, db: db}
}
