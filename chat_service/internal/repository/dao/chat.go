package dao

import (
	"errors"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
)

func (cd *chatDao) CreateChatHistory(chistory model.PrivateChatHistory) error {
	return cd.DB.Create(chistory).Error
}

func (cd *chatDao) GetChatHistories() ([]model.PrivateChatHistory, error) {
	return nil, errors.New("not implemented")
}

func (cd *chatDao) FindChatHistory() ([]model.PrivateChatHistory, error) {
	return nil, errors.New("not implemented")
}

func (cd *chatDao) RevokeChatHistory(id int64) error {
	return errors.New("not implemented")
}
