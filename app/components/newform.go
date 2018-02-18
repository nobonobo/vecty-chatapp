package components

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"

	bs4 "github.com/nobonobo/bootstrap4"
)

// NewForm ...
type NewForm struct {
	vecty.Core
	RoomName string `vecty:"prop"`
}

// Render ...
func (c *NewForm) Render() vecty.ComponentOrHTML {
	return &bs4.Card{
		Children: &bs4.CardBody{
			Children: vecty.List{
				elem.Form(
					&bs4.FormGroup{
						Children: vecty.List{
							&bs4.Label{For: "roomName", Children: vecty.Text("Room Name:")},
							&bs4.Input{
								Type:  prop.TypeText,
								Name:  "roomName",
								ID:    "roomName",
								Value: c.RoomName,
								Markup: vecty.Markup(
									event.Input(func(ev *vecty.Event) {
										c.RoomName = ev.Target.Get("value").String()
									}),
								),
							},
						},
					},
					&bs4.Button{
						Markup: vecty.Markup(event.Click(func(ev *vecty.Event) {
							log.Println("create new room:", c.RoomName)
							js.Global.Get("location").Set("hash", "#/room/"+c.RoomName)
						})),
						Children: vecty.Text("Create & Enter!"),
					},
				),
			},
		},
	}
}
