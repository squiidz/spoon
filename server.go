package main

import (
	"net/http"

	"local/spoon/api"

	"github.com/go-zoo/bone"
	"github.com/go-zoo/claw"
	mw "github.com/go-zoo/claw/middleware"
)

var (
	muxx   = bone.New()
	apiMux = bone.New()
	clw    = claw.New(mw.Logger)
)

func main() {
	muxx.SubRoute("/api", apiMux)
	muxx.GetFunc("/", api.IndexHandler)

	apiMux.GetFunc("/movie/:id", api.GetMovie)
	apiMux.GetFunc("/movie/title/:title", api.GetMovieByTitle)
	apiMux.PostFunc("/movie/new", api.PostMovie)

	http.ListenAndServe(":8080", clw.Merge(muxx))
}
