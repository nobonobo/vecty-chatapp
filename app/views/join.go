package views

import (
	"github.com/gopherjs/vecty"
	bs4 "github.com/nobonobo/bootstrap4"

	"github.com/nobonobo/vecty-sample/app/components"
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
		Contents: vecty.List{},
	}
}
