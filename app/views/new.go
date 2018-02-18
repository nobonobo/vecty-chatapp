package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	bs4 "github.com/nobonobo/bootstrap4"

	"github.com/nobonobo/vecty-sample/app/components"
)

// NewView ...
type NewView struct {
	vecty.Core
	RoomName string `vecty:"prop"`
}

// Render ...
func (c *NewView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("ChatApp:New")
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
			elem.Heading1(vecty.Text("New room")),
			&components.NewForm{RoomName: c.RoomName},
		},
	}
}
