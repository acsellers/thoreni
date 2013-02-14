package router

import "io"

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
	Render(string)
	io.Writer
}
