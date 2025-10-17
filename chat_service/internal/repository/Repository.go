package repository

type ChatRepository interface{}

type chatRepo struct {
}

func NewChatRepository() ChatRepository {
	return &chatRepo{}
}
