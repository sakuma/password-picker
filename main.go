package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {
	goji.Get("/", Root)
	goji.Serve()
}

func Root(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Password picker world!")
}
