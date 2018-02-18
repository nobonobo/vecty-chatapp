package views

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	bs4 "github.com/nobonobo/bootstrap4"

	"github.com/nobonobo/vecty-sample/app/components"
)

// RoomView ...
type RoomView struct {
	vecty.Core
	Name string
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
		},
	}
}
