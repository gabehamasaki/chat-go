package dtos

import (
	"github.com/gabehamasaki/chat-go/internal/database"
	"github.com/google/uuid"
)

type Chat struct {
	id         uuid.UUID `json:"id"`
	name       string    `json:"name"`
	created_at string    `json:"created_at"`
}

type ListChatsResponse struct {
	Chats []Chat `json:"chats"`
}

func ToChatsResponse(rawChats []database.Chat) ListChatsResponse {
	chats := make([]Chat, len(rawChats))
	for i, rawChat := range rawChats {
		chats[i] = Chat{
			id:         rawChat.ID,
			name:       rawChat.Name,
			created_at: rawChat.CreatedAt.Time.String(),
		}
	}
	return ListChatsResponse{Chats: chats}
}
