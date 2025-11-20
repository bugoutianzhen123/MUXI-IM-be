package repository

import "github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"

func (cr *chatRepo) CreateChatHistory(chistory model.PrivateChatHistory) error {
	return cr.CreateChatHistory(chistory)
}
func (cr *chatRepo) GetChatHistories() ([]model.PrivateChatHistory, error) {
	return cr.GetChatHistories()
}
func (cr *chatRepo) FindChatHistory() ([]model.PrivateChatHistory, error) {
	return cr.FindChatHistory()
}

func (cr *chatRepo) RevokeChatHistory(id int64) error {
	return cr.RevokeChatHistory(id)
}
