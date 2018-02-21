package backend

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"path"

	"golang.org/x/net/websocket"

	"github.com/google/uuid"
	"github.com/nobonobo/vecty-sample/app/models"
)

// Setup ...
func Setup(ctx context.Context, prefix string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(prefix+"/new", newRoom)
	mux.Handle(prefix+"/join/", websocket.Handler(joinRoom))
	go roomManage(ctx)
	return mux
}

func newRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	room := &models.Room{
		UUID: uuid.New(),
	}
	newRoomCh <- room
	res := &models.NewRoomRes{
		RoomID: room.UUID,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func joinRoom(ws *websocket.Conn) {
	_, f := path.Split(ws.Request().URL.Path)
	room := GetRoom(f)
	if room == nil {
		log.Println("unknown room id:", f)
		ws.Close()
		return
	}
	var member *models.Member
	if err := json.NewDecoder(ws).Decode(&member); err != nil {
		log.Println(err)
		ws.Close()
		return
	}
	ctx := ws.Request().Context()
	m := NewMember(ctx, room, member, ws)
	for _, member := range room.GetMembers() {
		if err := m.Write("join", member); err != nil {
			log.Println(err)
		}
	}
	room.Join(m)
	<-ctx.Done()
}
