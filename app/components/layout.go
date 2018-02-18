package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Layout ...
type Layout struct {
	vecty.Core
	MenuItems vecty.ComponentOrHTML `vecty:"prop"`
	Contents  vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Layout) Render() vecty.ComponentOrHTML {
	return elem.Body(
		&Navbar{
			MenuItems: c.MenuItems,
		},
		elem.Div(
			vecty.Markup(vecty.Class("container")),
			c.Contents,
		),
	)
}
