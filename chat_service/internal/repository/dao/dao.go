package dao

type ChatDao interface {
	CreateChatHistory() error
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
