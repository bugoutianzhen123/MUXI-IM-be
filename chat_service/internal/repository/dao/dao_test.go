package dao

import (
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateChatHistory(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	cd := NewDao(db)
	if err := cd.CreateChatHistory(model.PrivateChatHistory{}); err != nil {
		t.Fatal(err)
	}
	gcd := NewGroupChatDao()
	if err := gcd.CreateChatHistory(model.GroupChatHistory{}); err != nil {
		t.Fatal(err)
	}
}
