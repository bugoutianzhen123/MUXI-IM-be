package dao

import "testing"

func TestCreateChatHistory(t *testing.T) {
	gcd := NewGroupChatDao()
	if err := gcd.CreateChatHistory(); err != nil {
		t.Fatal(err)
	}
}
