package router

import (
	"log"

	orig "github.com/go-humble/router"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

var (
	r           = orig.New()
	currentView vecty.Component
)

func init() {
	r.ForceHashURL = true
	js.Global.Set("rerender", Rerender)
}

// HandleFunc ...
func HandleFunc(path string, handler orig.Handler) {
	r.HandleFunc(path, handler)
}

// Navigate ...
func Navigate(uri string) {
	r.Navigate(uri)
}

// Start ...
func Start() {
	r.Start()
}

// Stop ...
func Stop() {
	r.Stop()
}

// RenderBody ...
func RenderBody(body vecty.Component) {
	if v, ok := currentView.(Teardowner); ok {
		v.Teardown()
	}
	currentView = body
	js.Global.Set("currentView", body)
	vecty.RenderBody(currentView)
	if v, ok := body.(Setuper); ok {
		v.Setup()
	}
}

// Rerender ...
func Rerender() {
	if currentView != nil {
		log.Println("rerender:", currentView)
		vecty.Rerender(currentView)
	}
}

// CurrentView ...
func CurrentView() vecty.Component {
	return currentView
}
