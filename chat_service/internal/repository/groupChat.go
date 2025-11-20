package repository

import "github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"

func (gr *groupRepo) CreateChatHistory(gchistory model.GroupChatHistory) error {
	return gr.CreateChatHistory(gchistory)
}
func (gr *groupRepo) GetChatHistories() ([]model.GroupChatHistory, error) {
	return gr.GetChatHistories()
}
func (gr *groupRepo) FindChatHistory() ([]model.GroupChatHistory, error) {
	return gr.FindChatHistory()
}
