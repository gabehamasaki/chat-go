package database

import (
	"context"
	"fmt"

	"github.com/gabehamasaki/chat-go/internal/config"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func NewDatabase(logger *zap.Logger, cfg *config.Config) *Queries {
	ctx := context.Background()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		logger.Error("Error on connect to database", zap.Error(err))
		return nil
	}

	logger.Info("Connected to database")
	queries := New(conn)

	return queries
}
