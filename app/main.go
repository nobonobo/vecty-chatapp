package main

import (
	orig "github.com/go-humble/router"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"

	_ "github.com/nobonobo/vecty-sample/app/handlers"
	"github.com/nobonobo/vecty-sample/app/router"
	_ "github.com/nobonobo/vecty-sample/app/store"
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

	router.HandleFunc("/", func(ctx *orig.Context) {
		router.RenderBody(&views.TopView{})
	})
	router.HandleFunc("/new", func(ctx *orig.Context) {
		router.RenderBody(&views.NewView{})
	})
	router.HandleFunc("/join/{uuid}", func(ctx *orig.Context) {
		router.RenderBody(&views.JoinView{Name: ctx.Params["uuid"]})
	})
	router.HandleFunc("/room/{uuid}", func(ctx *orig.Context) {
		router.RenderBody(&views.RoomView{Name: ctx.Params["uuid"]})
	})
	router.Start()
}
