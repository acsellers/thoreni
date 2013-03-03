package app

import (
	"github.com/acsellers/thoreni"
	"github.com/acsellers/thoreni/logging"
	"github.com/acsellers/thoreni/render"
	"github.com/acsellers/thoreni/router"
	"net/http"
)

type RunnableApp struct {
	Routing   Routable
	Templates thoreni.Renderable
}

type Routable interface {
	Match(req *http.Request) router.RoutingFunc
}

func (r RunnableApp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	renderer := render.NewRequestRenderer()
	renderer.Output = w
	renderer.Request = req

	miniLogger := NewMiniLogger()
	defer miniLogger.Flush()

	miniLogger.LogRequest(req)

	contextable := &thoreni.Contextable{LocalRenderer: renderer, Request: req, Logger, miniLogger}
	r.Routing.Match(req)(contextable)
	miniLogger.CloseRequest(contextable)
}

func ListenAndServe(user_app *RunnableApp, port string) error {
	http.DefaultServeMux = http.NewServeMux()
	http.Handle("/", user_app)
	return http.ListenAndServe(port, nil)
}
