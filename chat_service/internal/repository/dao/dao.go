package dao

import "gorm.io/gorm"

type ChatDao interface {
	CreateChatHistory() error
	GetChatHistory() error
	FindChatHistory() error
}

type GroupChatDao interface {
	ChatDao
}

type chatDao struct {
}

type groupChatDao struct {
	chatDao
}

func NewDao() ChatDao {
	return &chatDao{}
}

func NewGroupChatDao() GroupChatDao {
	return &groupChatDao{}
}

func InitTables(db *gorm.DB) error {
	return nil
}
