package dao

import (
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
	"gorm.io/gorm"
)

type ChatDao interface {
	CreateChatHistory(chistory model.PrivateChatHistory) error
	GetChatHistories() ([]model.PrivateChatHistory, error)
	FindChatHistory() ([]model.PrivateChatHistory, error)
}

type GroupChatDao interface {
	CreateChatHistory(gchistory model.GroupChatHistory) error
	GetChatHistories() ([]model.GroupChatHistory, error)
	FindChatHistory() ([]model.GroupChatHistory, error)
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
	return nil
}
