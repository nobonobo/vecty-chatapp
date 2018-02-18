package main

import (
	"github.com/go-humble/router"
	"github.com/google/uuid"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"

	"github.com/nobonobo/vecty-sample/app/views"
)

func main() {
	document := js.Global.Get("document")
	meta := document.Call("createElement", "meta")
	meta.Set("name", "viewport")
	meta.Set("content", "width=device-width, initial-scale=1, shrink-to-fit=no")
	document.Get("head").Call("appendChild", meta)
	vecty.AddStylesheet("https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("assets/app.css")

	r := router.New()
	r.ForceHashURL = true
	r.HandleFunc("/", func(ctx *router.Context) {
		vecty.RenderBody(&views.TopView{})
	})
	r.HandleFunc("/new", func(ctx *router.Context) {
		vecty.RenderBody(&views.NewView{
			RoomName: uuid.New().String(),
		})
	})
	r.HandleFunc("/join/{name}", func(ctx *router.Context) {
		vecty.RenderBody(&views.JoinView{Name: ctx.Params["name"]})
	})
	r.HandleFunc("/room/{name}", func(ctx *router.Context) {
		vecty.RenderBody(&views.RoomView{Name: ctx.Params["name"]})
	})
	r.Start()
}
