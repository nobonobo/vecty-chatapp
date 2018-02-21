package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/nobonobo/vecty-sample/backend"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var development bool
	flag.BoolVar(&development, "dev", false, "reverseproxy to gopherjs serve(localhost:8080)")
	flag.Parse()
	if development {
		log.Println("development mode")
		u, _ := url.Parse("http://localhost:8080")
		rp := httputil.NewSingleHostReverseProxy(u)
		http.Handle("/", rp)
	} else {
		log.Println("normal mode")
		http.Handle("/", http.FileServer(http.Dir("./app")))
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	h := backend.Setup(ctx, "/api")
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("listen:", l.Addr())
	if err := http.Serve(l, nil); err != nil {
		log.Fatalln(err)
	}
}
