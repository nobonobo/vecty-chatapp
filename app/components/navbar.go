package components

import (
	"github.com/gopherjs/vecty"
	bs4 "github.com/nobonobo/bootstrap4"
)

// Navbar ...
type Navbar struct {
	vecty.Core
	MenuItems vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Navbar) Render() vecty.ComponentOrHTML {
	return &bs4.Navbar{
		Size:     bs4.SizeLarge,
		Expand:   true,
		FixedTop: true,
		Light:    true,
		Children: vecty.List{
			&bs4.NavbarBrand{Href: "#/", Children: vecty.Text("CahtApp")},
			&bs4.NavbarToggler{Type: "button", Target: "#topBarMenues"},
			&bs4.NavbarCollapse{
				ID: "topBarMenues",
				Children: &bs4.NavbarNav{
					Children: c.MenuItems,
				},
			},
		},
	}
}
