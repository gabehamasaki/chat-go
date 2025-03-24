package handler

import (
	"net/http"

	"github.com/gabehamasaki/chat-go/internal/dtos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) ListChats(ctx *gin.Context) {
	rawChats, err := h.db.GetAllChats(ctx)
	if err != nil {
		h.logger.Error("Failed to get all chats", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all chats"})
		return
	}

	chats := dtos.ToChatsResponse(rawChats)

	ctx.JSON(http.StatusOK, chats)
}
