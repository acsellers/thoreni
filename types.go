package thoreni

import (
	"net/http"
)

type Contextable struct {
	Renderable
	Request *http.Request
	Requestish
}

type Requestish interface {
	Path() string
	Method() string
}
type Renderable interface {
	Render(string, interface{})
	RenderStatic(string, interface{})
	Layout(string)
	Redirect(address string)
}
