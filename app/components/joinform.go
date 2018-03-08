package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"

	bs4 "github.com/nobonobo/bootstrap4"
	"github.com/nobonobo/vecty-chatapp/app/actions"
	"github.com/nobonobo/vecty-chatapp/app/dispatcher"
)

// JoinForm ...
type JoinForm struct {
	vecty.Core
	Label    string `vecty:"prop"`
	New      bool   `vecty:"prop"`
	RoomID   string `vecty:"prop"`
	Nickname string `vecty:"prop"`
}

// Render ...
func (c *JoinForm) Render() vecty.ComponentOrHTML {
	return &bs4.Card{
		Children: &bs4.CardBody{
			Children: vecty.List{
				elem.Form(
					&bs4.FormGroup{
						Children: vecty.List{
							&bs4.Label{For: "nickname", Children: vecty.Text("Nick-Name:")},
							&bs4.Input{
								Type:  prop.TypeText,
								Name:  "nickname",
								ID:    "nickname",
								Value: c.Nickname,
								Markup: vecty.Markup(
									event.Input(func(ev *vecty.Event) {
										c.Nickname = ev.Target.Get("value").String()
									}),
								),
							},
						},
					},
					&bs4.Button{
						Type: prop.TypeSubmit,
						Markup: vecty.Markup(event.Click(func(ev *vecty.Event) {
							if c.New {
								dispatcher.Dispatch(actions.NewRoom{
									Nickname: c.Nickname,
								})
							} else {
								dispatcher.Dispatch(actions.JoinRoom{
									RoomID:   c.RoomID,
									Nickname: c.Nickname,
								})
							}
						}).PreventDefault()),
						Children: vecty.Text(c.Label),
					},
				),
			},
		},
	}
}
