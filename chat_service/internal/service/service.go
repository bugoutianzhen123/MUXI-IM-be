package service

type ChatService interface {
	SendMessage(message string)
	GetHistoryMessage() string
	SearchMessage(message string)
	RevokeMessage()
	// 消息状态
    MarkMessageAsRead()
    GetUnreadCount()

	UpdateUserStatus()
	GetUserStatus()

	SubscribeMessages()
}

type GroupChatService interface {
	ChatService
	// group
	CreateGroup()
	GetSelfGroups()
	JoinGroup()
	LeaveGroup()
	DismissGroup()
	// member
	GetGroupMembers()
	AddMember()
	RemoveMember()
	UpdateMemberRole()

	UpdateGroupInfo()
	GetGroupInfo()
}
