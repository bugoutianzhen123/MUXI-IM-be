package service

import "github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"

type ChatService interface {
	Create(history model.PrivateChatHistory) error
	Get() ([]model.PrivateChatHistory, error)
	Find() ([]model.PrivateChatHistory, error)
	Revoke() error
	// 消息状态
	//MarkMessageAsRead()
	//GetUnreadCount()
	//
	//UpdateUserStatus()
	//GetUserStatus()
	//
	//SubscribeMessages()
}
