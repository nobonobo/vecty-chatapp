package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/url"
	"time"

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
	Name     string            `vecty:"prop"`
	Members  []*models.Member  `vecty:"prop"`
	Messages []*models.Message `vecty:"prop"`
	conn     net.Conn
	encoder  *json.Encoder
}

// Render ...
func (c *RoomView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("ChatApp:Room")
	loc := js.Global.Get("location")
	base := loc.Get("origin").String() + loc.Get("pathname").String()
	href := fmt.Sprintf("%s#/join/%s", base, c.Name)
	members := vecty.List{}
	for _, m := range c.Members {
		members = append(members, elem.Div(
			vecty.Markup(vecty.ClassMap{
				"alert":      true,
				"alert-info": true,
			}),
			vecty.Text(m.Nickname),
		))
	}
	messages := vecty.List{}
	for _, m := range c.Messages {
		messages = append(messages, &bs4.Card{
			Children: &bs4.CardBody{
				Children: vecty.Text(fmt.Sprintf("%s : %s", m.Nickname, m.Content)),
			},
		})
	}
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
		Contents: elem.Div(
			vecty.Markup(vecty.Class("row")),
			elem.Div(
				vecty.Markup(vecty.ClassMap{
					"col-12":   true,
					"col-md-8": true,
					"col-lg-9": true,
					"col-xl-9": true,
				}),
				elem.Heading5(vecty.Text("Room:"+c.Name)),
				messages,
				&components.ChatForm{Markup: vecty.Markup(prop.ID("message"))},
			),
			elem.Div(
				vecty.Markup(vecty.ClassMap{
					"col-12":   true,
					"col-md-4": true,
					"col-lg-3": true,
					"col-xl-3": true,
				}),
				elem.Heading5(vecty.Text("Join Link")),
				elem.Anchor(
					vecty.Markup(prop.Href(href), vecty.Attribute("target", "_blank")),
					&components.QRCode{Text: href, CellSize: 5},
				),
				elem.Heading5(vecty.Text("Members")),
				&bs4.Card{
					Children: &bs4.CardBody{
						Children: members,
					},
				},
			),
		),
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

// ShowForm ...
func ShowForm() {
	js.Global.Get("document").
		Call("querySelector", "#message").
		Call("scrollIntoViewIfNeeded")
}

// Setup ...
func (c *RoomView) Setup() {
	js.Global.Set("showForm", ShowForm)
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
				c.Messages = append(c.Messages, message)
				if len(c.Members) > 100 {
					c.Messages = c.Messages[:100]
				}
				vecty.Rerender(c)
				time.AfterFunc(200*time.Millisecond, ShowForm)
			case "join":
				var member *models.Member
				if err := v.Unmarshal(&member); err != nil {
					log.Println(err)
				}
				log.Println("join", member)
				c.Members = append(c.Members, member)
				vecty.Rerender(c)
			case "leave":
				var member *models.Member
				if err := v.Unmarshal(&member); err != nil {
					log.Println(err)
				}
				log.Println("leave", member)
				for i, m := range c.Members {
					if member.UUID == m.UUID {
						c.Members = append(c.Members[:i], c.Members[i+1:]...)
						vecty.Rerender(c)
						break
					}
				}
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
