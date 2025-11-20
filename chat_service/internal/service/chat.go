package service

import (
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
)

func (cs *chatService) Create(history model.PrivateChatHistory) error {
	return cs.repo.CreateChatHistory(history)
}

func (cs *chatService) Get() ([]model.PrivateChatHistory, error) {
	return cs.repo.GetChatHistories()
}

func (cs *chatService) Find() ([]model.PrivateChatHistory, error) {
	return cs.repo.FindChatHistory()
}

func (cs *chatService) Revoke() error {
	return cs.repo.RevokeChatHistory(-1)
}
