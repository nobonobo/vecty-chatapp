package backend

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/nobonobo/vecty-sample/app/models"
	"golang.org/x/net/websocket"
)

// Member ...
type Member struct {
	UUID     uuid.UUID
	Nickname string
	parent   *Room
	conn     *websocket.Conn
	encoder  *json.Encoder
	cancel   func()
	once     sync.Once
}

// NewMember ...
func NewMember(ctx context.Context, parent *Room, member *models.Member, conn *websocket.Conn) *Member {
	m := &Member{
		UUID:     member.UUID,
		Nickname: member.Nickname,
		parent:   parent,
		conn:     conn,
		encoder:  json.NewEncoder(conn),
	}
	ctx, m.cancel = context.WithCancel(ctx)
	go m.do(ctx)
	return m
}

func (m *Member) do(ctx context.Context) {
	defer m.Close()
	defer m.parent.Leave(m.UUID)
	reader := json.NewDecoder(m.conn)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		var v *models.Event
		if err := reader.Decode(&v); err != nil {
			if err == io.EOF {
				return
			}
			log.Println(err)
			continue
		}
		log.Println("received:", v.Type, string(v.Data))
		switch v.Type {
		case "message":
			var message *models.Message
			if err := v.Unmarshal(&message); err != nil {
				log.Println(err)
				continue
			}
			m.parent.Publish(message)
		}
	}
}

// Close ...
func (m *Member) Close() {
	m.once.Do(func() {
		var conn *websocket.Conn
		conn, m.conn = m.conn, nil
		conn.Close()
		m.cancel()
	})
}

// Write ...
func (m *Member) Write(tp string, obj interface{}) error {
	return m.encoder.Encode(struct {
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	}{tp, obj})
}
