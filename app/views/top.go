package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	bs4 "github.com/nobonobo/bootstrap4"

	"github.com/nobonobo/vecty-chatapp/app/components"
)

// TopView ...
type TopView struct {
	vecty.Core
}

// Render ...
func (c *TopView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("ChatApp:Top")
	return &components.Layout{
		MenuItems: vecty.List{
			&bs4.NavItem{
				Active: true,
				Children: &bs4.NavLink{
					Href:     "#/",
					Children: vecty.Text("Top"),
				},
			},
			&bs4.NavItem{
				Active: false,
				Children: &bs4.NavLink{
					Href:     "#/new",
					Children: vecty.Text("New"),
				},
			},
		},
		Contents: &bs4.Jumbotron{
			Children: vecty.List{
				elem.Heading1(vecty.Text("Top")),
				elem.Paragraph(
					vecty.Markup(vecty.Class("lead")),
					vecty.Text("dead-simple chat app"),
				),
				elem.Paragraph(
					vecty.Markup(vecty.Class("lead")),
					&bs4.ButtonLinks{
						Size:     bs4.SizeLarge,
						Href:     "#/new",
						Children: vecty.Text("Create New Room"),
					},
				),
			},
		},
	}
}
