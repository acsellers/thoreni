package thoreni

import (
	"github.com/acsellers/thoreni/logging"
	"net/http"
)

type Contextable struct {
	LocalRenderer
	Request *http.Request
	Logger  *logging.MiniLogger
}

type LocalRenderer interface {
	Render(string)
	RenderStatic(string)
	Layout(string)
	Redirect(address string)
}

type Renderable interface {
	RenderTemplate(string, string, interface{}) (string, string, string)
	RenderStatic(string, string, interface{}) (string, string, string)
}
