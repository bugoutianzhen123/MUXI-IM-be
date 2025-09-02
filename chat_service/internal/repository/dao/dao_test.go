package dao

import "testing"

func TestCreateChatHistory(t *testing.T) {
	cd := NewDao()
	if err := cd.CreateChatHistory(); err != nil {
		t.Fatal(err)
	}
	gcd := NewGroupChatDao()
	if err := gcd.CreateChatHistory(); err != nil {
		t.Fatal(err)
	}
}
