package app

import (
	"github.com/acsellers/thoreni/render"
	"github.com/acsellers/thoreni/router"
)

type RunnableApp struct {
	Routing   Routable
	Templates Templateable
}

type Routable interface {
	Match(req Requestish) RoutingFunc
}

type Templateable interface {
	RenderTemplate(string, string, interface{}) (string, string, string)
	RenderStatic(string, string, interface{}) (string, string, string)
	Layout(string)
	Redirect(string)
}

func Run(user_app *Runnable) {

}
