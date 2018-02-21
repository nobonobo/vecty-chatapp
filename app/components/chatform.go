package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"

	bs4 "github.com/nobonobo/bootstrap4"
	"github.com/nobonobo/vecty-sample/app/actions"
	"github.com/nobonobo/vecty-sample/app/dispatcher"
)

// ChatForm ...
type ChatForm struct {
	vecty.Core
	Message string `vecty:"prop"`
}

// Render ...
func (c *ChatForm) Render() vecty.ComponentOrHTML {
	return &bs4.Card{
		Children: &bs4.CardBody{
			Children: vecty.List{
				elem.Form(
					&bs4.FormGroup{
						Children: vecty.List{
							&bs4.Label{For: "message", Children: vecty.Text("Message:")},
							&bs4.Input{
								Type:  prop.TypeText,
								Name:  "message",
								ID:    "message",
								Value: c.Message,
								Markup: vecty.Markup(
									event.Input(func(ev *vecty.Event) {
										c.Message = ev.Target.Get("value").String()
									}),
								),
							},
						},
					},
					&bs4.Button{
						Type: prop.TypeSubmit,
						Markup: vecty.Markup(event.Click(func(ev *vecty.Event) {
							dispatcher.Dispatch(actions.Message{
								Message: c.Message,
							})
							c.Message = ""
							vecty.Rerender(c)
						}).PreventDefault()),
						Children: vecty.Text("Send"),
					},
				),
			},
		},
	}
}
