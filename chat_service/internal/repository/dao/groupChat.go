package dao

import (
	"errors"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
)

func (gcd *groupChatDao) CreateGroupChatHistory(gchistory model.GroupChatHistory) error {
	return errors.New("not implemented")
}

func (gcd *groupChatDao) GetGroupChatHistories() ([]model.GroupChatHistory, error) {
	return nil, errors.New("not implemented")
}

func (gcd *groupChatDao) FindGroupChatHistory() ([]model.GroupChatHistory, error) {
	return nil, errors.New("not implemented")
}
