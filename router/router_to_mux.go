package router

import (
	"fmt"
	"net/http"
)

// GoNative will return a ServeMux that you could use for serving on a specific port, or any number 
// of interesting things.
func (r *Router) GoNative() *http.ServeMux {
	sm := http.NewServeMux()
	sm.Handle("/", r)
	return sm
}

// TakeOver will force this Router to overrule any other route definitions on the DefaultServeMux,
// it will reset the DefaultServeMux before doing this to ensure that it wins.
func (r *Router) TakeOver() {
	http.DefaultServeMux = http.NewServeMux()
	http.Handle("/", r)
}

// ServeHTTP is here so that a thoreni.Router can be used as a valid Handler for the net/http
// system. You are not expected to call this function yourself, but you can if you want.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	wrappedRequest := wrapRequest(req)
	fixmeContext := &Contextable{Renderable: &FIXMERenderer{w}, Requestish: wrappedRequest}
	r.Match(wrappedRequest)(fixmeContext)
}

// A thoreni.Request makes it a bit easier to access the Path from the request, and eventually 
// it will allow you to carry request related data from hooks to RoutingFuncs
type Request struct {
	Req *http.Request
}

func (vr Request) Method() string {
	return vr.Req.Method
}

func (vr Request) Path() string {
	return vr.Req.URL.Path
}

func wrapRequest(req *http.Request) *Request {
	return &Request{req}
}

// FIXMERenderer, it does nothing, it just wraps a http.ResponseWriter and writes to it. Eventually
// I will write an actual Renderer class that does all sorts of useful things
type FIXMERenderer struct {
	w http.ResponseWriter
}

func (fr *FIXMERenderer) Render(s string) {
	fmt.Fprintf(fr.w, s)
}

func (fr *FIXMERenderer) Write(p []byte) (n int, err error) {
	return fr.w.Write(p)
}
