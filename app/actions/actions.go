package actions

// NewRoom ...
type NewRoom struct {
	base
	Nickname string
}

// JoinRoom ...
type JoinRoom struct {
	base
	RoomID   string
	Nickname string
}

type Message struct {
	base
	Message string
}
