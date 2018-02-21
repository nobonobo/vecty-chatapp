package models

import "github.com/google/uuid"

// Message ...
type Message struct {
	Author   uuid.UUID
	Nickname string
	Content  string
}

// Member ...
type Member struct {
	UUID     uuid.UUID
	Nickname string
}

// Room ...
type Room struct {
	UUID    uuid.UUID
	Members []Member
	History []Message
}
