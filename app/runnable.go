package app

import (
	"github.com/acsellers/thoreni/render"
	"github.com/acsellers/thoreni/router"
)

type Runnable struct {
	router.Routable
	render.Templateable
	render.Staticable
}

func Run(user_app Runnable) {

}
