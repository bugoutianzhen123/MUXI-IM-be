package dao

import "github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"

func (cd *chatDao) CreateChatHistory(chistory model.PrivateChatHistory) error {
	return cd.DB.Create(chistory).Error
}

func (cd *chatDao) GetChatHistories() ([]model.PrivateChatHistory, error) {
	return nil, nil
}

func (cd *chatDao) FindChatHistory() ([]model.PrivateChatHistory, error) {
	return nil, nil
}
