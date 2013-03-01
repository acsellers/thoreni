package router

type RoutingFunc func(*Contextable)

type Contextable struct {
	Renderable
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
