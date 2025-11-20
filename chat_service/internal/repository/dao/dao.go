package dao

import (
	"errors"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
	"gorm.io/gorm"
)

type ChatDao interface {
	CreateChatHistory(chistory model.PrivateChatHistory) error
	GetChatHistories() ([]model.PrivateChatHistory, error)
	FindChatHistory() ([]model.PrivateChatHistory, error)
	RevokeChatHistory(id int64) error
}

type GroupChatDao interface {
	CreateGroupChatHistory(gchistory model.GroupChatHistory) error
	GetGroupChatHistories() ([]model.GroupChatHistory, error)
	FindGroupChatHistory() ([]model.GroupChatHistory, error)
}

type chatDao struct {
	*gorm.DB
}

type groupChatDao struct {
	chatDao
}

func NewDao(db *gorm.DB) ChatDao {
	return &chatDao{db}
}

func NewGroupChatDao() GroupChatDao {
	return &groupChatDao{}
}

func InitTables(db *gorm.DB) error {
	return errors.New("not implemented")
}
