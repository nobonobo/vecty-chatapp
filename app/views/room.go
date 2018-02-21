package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/goxjs/websocket"
	bs4 "github.com/nobonobo/bootstrap4"

	"github.com/nobonobo/vecty-sample/app/components"
	"github.com/nobonobo/vecty-sample/app/models"
	"github.com/nobonobo/vecty-sample/app/router"
	"github.com/nobonobo/vecty-sample/app/store"
)

// RoomView ...
type RoomView struct {
	vecty.Core
	Name    string `vecty:"prop"`
	conn    net.Conn
	encoder *json.Encoder
}

// Render ...
func (c *RoomView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("ChatApp:Room")
	loc := js.Global.Get("location")
	base := loc.Get("origin").String() + loc.Get("pathname").String()
	href := fmt.Sprintf("%s#/join/%s", base, c.Name)
	return &components.Layout{
		MenuItems: vecty.List{
			&bs4.NavItem{
				Active: false,
				Children: &bs4.NavLink{
					Href:     "#/",
					Children: vecty.Text("Top"),
				},
			},
		},
		Contents: vecty.List{
			elem.Heading4(vecty.Text("Room:" + c.Name)),
			&components.QRCode{Text: href, CellSize: 5},
			elem.Anchor(
				vecty.Markup(prop.Href(href), vecty.Attribute("target", "_blank")),
				vecty.Text("open link"),
			),
			&components.ChatForm{},
		},
	}
}

// Publish ...
func (c *RoomView) Publish(message string) error {
	m := &models.Message{
		Author:   store.UUID,
		Nickname: store.Nickname,
		Content:  message,
	}
	log.Println("publish:", m)
	return c.encoder.Encode(map[string]interface{}{
		"type": "message",
		"data": m,
	})
}

// Setup ...
func (c *RoomView) Setup() {
	log.Println("setup:", c.Name)
	origin := js.Global.Get("location").Get("origin").String()
	u, err := url.Parse(origin)
	if err != nil {
		log.Println(err)
		return
	}
	switch u.Scheme {
	case "http":
		u.Scheme = "ws"
	case "https":
		u.Scheme = "wss"
	}
	u.Path = "/api/join/" + c.Name
	conn, err := websocket.Dial(u.String(), origin)
	if err != nil {
		log.Println("ws connect failed:", err)
		router.Navigate("/")
		return
	}
	log.Println("ws connected:", u.String())
	c.conn = conn
	c.encoder = json.NewEncoder(conn)
	go func() {
		decoder := json.NewDecoder(c.conn)
		for {
			var v *models.Event
			if err := decoder.Decode(&v); err != nil {
				log.Println("ws error:", err)
				router.Navigate("/")
				return
			}
			switch v.Type {
			case "message":
				var message *models.Message
				if err := v.Unmarshal(&message); err != nil {
					log.Println(err)
				}
				log.Println("message:", message)
			case "join":
				var member *models.Member
				if err := v.Unmarshal(&member); err != nil {
					log.Println(err)
				}
				log.Println("join", member)
			case "leave":
				var member *models.Member
				if err := v.Unmarshal(&member); err != nil {
					log.Println(err)
				}
				log.Println("leave", member)
			}
		}
	}()
	member := &models.Member{UUID: store.UUID, Nickname: store.Nickname}
	if err := c.encoder.Encode(member); err != nil {
		log.Println("join failed:", err)
		router.Navigate("/")
	}
}

// Teardown ...
func (c *RoomView) Teardown() {
	c.conn.Close()
}
