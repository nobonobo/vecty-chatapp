package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

// NewRoomRes ...
type NewRoomRes struct {
	RoomID uuid.UUID `json:"roomId"`
}

// Event ...
type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// Unmarshal ...
func (e *Event) Unmarshal(v interface{}) error {
	return json.Unmarshal(e.Data, v)
}
