package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nobonobo/vecty-sample/app/models"

	"github.com/nobonobo/vecty-sample/app/actions"
	"github.com/nobonobo/vecty-sample/app/dispatcher"
	"github.com/nobonobo/vecty-sample/app/router"
	"github.com/nobonobo/vecty-sample/app/store"
)

func init() {
	dispatcher.Register(handler)
}

// Publisher ...
type Publisher interface {
	Publish(message string) error
}

func handler(a actions.Action) {
	log.Println("handle action:", a)
	switch act := a.(type) {
	case actions.NewRoom:
		go newRoom(act)
	case actions.JoinRoom:
		store.Nickname = act.Nickname
		router.Navigate("/room/" + act.RoomID)
	case actions.Message:
		if p, ok := router.CurrentView().(Publisher); ok {
			if err := p.Publish(act.Message); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func newRoom(act actions.NewRoom) {
	resp, err := http.Post("/api/new", "application/json", nil)
	if err != nil {
		log.Println(err)
		return
	}
	var v models.NewRoomRes
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		log.Println(err)
		return
	}
	store.Nickname = act.Nickname
	router.Navigate("/room/" + v.RoomID.String())
}
