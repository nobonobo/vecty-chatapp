package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	bs4 "github.com/nobonobo/bootstrap4"

	"github.com/nobonobo/vecty-chatapp/app/components"
	"github.com/nobonobo/vecty-chatapp/app/store"
)

// NewView ...
type NewView struct {
	vecty.Core
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
			elem.Heading1(vecty.Text("New Room")),
			&components.JoinForm{New: true, Label: "Create & Enter", Nickname: store.Nickname},
		},
	}
}
