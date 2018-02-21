package backend

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nobonobo/vecty-sample/app/models"
)

// TIMEOUT ...
const TIMEOUT = 300 * time.Second

type getRoom struct {
	uuid   string
	result chan *Room
}

var (
	newRoomCh = make(chan *models.Room)
	delRoomCh = make(chan uuid.UUID, 1)
	getRoomCh = make(chan getRoom)
)

// GetRoom ...
func GetRoom(uid string) *Room {
	ch := make(chan *Room, 1)
	getRoomCh <- getRoom{uid, ch}
	return <-ch
}

func roomManage(ctx context.Context) {
	rooms := map[string]*Room{}
	for {
		select {
		case <-ctx.Done():
			return
		case u := <-delRoomCh:
			rooms[u.String()].Close()
			delete(rooms, u.String())
			log.Println("del room:", u.String())
		case room := <-newRoomCh:
			rooms[room.UUID.String()] = NewRoom(context.Background(), room)
			log.Println("new room:", room.UUID.String())
		case req := <-getRoomCh:
			req.result <- rooms[req.uuid]
		}
	}
}

// Room ...
type Room struct {
	UUID    uuid.UUID
	joinCh  chan *Member
	leaveCh chan uuid.UUID
	msgCh   chan *models.Message
	timer   *time.Timer
	cancel  func()
	once    sync.Once
}

// NewRoom ...
func NewRoom(ctx context.Context, room *models.Room) *Room {
	r := &Room{
		UUID:    room.UUID,
		joinCh:  make(chan *Member),
		leaveCh: make(chan uuid.UUID),
		msgCh:   make(chan *models.Message),
		timer: time.AfterFunc(TIMEOUT, func() {
			delRoomCh <- room.UUID
		}),
	}
	ctx, r.cancel = context.WithCancel(ctx)
	go r.do(ctx)
	return r
}

func (r *Room) do(ctx context.Context) {
	log.Println("start do room:", r.UUID)
	defer log.Println("stop do room:", r.UUID)
	members := map[string]*Member{}
	defer func() {
		for _, m := range members {
			m.Close()
		}
		r.Close()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case member := <-r.joinCh:
			join := &models.Member{
				UUID:     member.UUID,
				Nickname: member.Nickname,
			}
			members[member.UUID.String()] = member
			for _, m := range members {
				if err := m.Write("join", join); err != nil {
					log.Println(err)
				}
			}
		case u := <-r.leaveCh:
			member := members[u.String()]
			if member != nil {
				leave := &models.Member{
					UUID:     member.UUID,
					Nickname: member.Nickname,
				}
				for _, m := range members {
					if err := m.Write("leave", leave); err != nil {
						log.Println(err)
					}
				}
				member.Close()
			}
			delete(members, u.String())
		case msg := <-r.msgCh:
			for _, m := range members {
				if err := m.Write("message", msg); err != nil {
					log.Println(err)
				}
			}
		}
		r.timer.Reset(TIMEOUT)
	}
}

// Close ...
func (r *Room) Close() {
	r.once.Do(func() {
		r.cancel()
	})
}

// Publish ...
func (r *Room) Publish(msg *models.Message) {
	r.msgCh <- msg
}

// Join ...
func (r *Room) Join(member *Member) {
	r.joinCh <- member
}

// Leave ...
func (r *Room) Leave(u uuid.UUID) {
	r.leaveCh <- u
}
