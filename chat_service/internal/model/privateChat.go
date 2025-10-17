package model

import "time"

type PrivateChatHistory struct {
	Id           int64
	Sender_uid   int64
	Receiver_uid int64
	QuoteReply   int64
	Content      []MessageContent
	Timestamp    int64
	Status       int               // 消息状态
	Extensions   map[string]string // 可能是额外的元数据

	CreatedAt time.Time
	UpdatedAt time.Time
}

type PrivateChatMessage struct {
	Id           int64
	Sender_uid   int64
	Receiver_uid int64
	QuoteReply   int64
	Content      []MessageContent
	Timestamp    int64
	Status       int               // 消息状态
	Extensions   map[string]string // 可能是额外的元数据
}

type MessageContent struct {
	Kind    int // At, PlainText, Image
	Content string
}
