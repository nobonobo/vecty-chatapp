package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	bs4 "github.com/nobonobo/bootstrap4"

	"github.com/nobonobo/vecty-sample/app/components"
	"github.com/nobonobo/vecty-sample/app/store"
)

// JoinView ...
type JoinView struct {
	vecty.Core
	Name string
}

// Render ...
func (c *JoinView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("ChatApp:Join")
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
			elem.Heading1(vecty.Text("Join Room")),
			&components.JoinForm{Label: "Join", RoomID: c.Name, Nickname: store.Nickname},
		},
	}
}
