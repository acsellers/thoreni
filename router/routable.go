package router

const (
	GET     = "G"
	POST    = "P"
	PUT     = "T"
	UPDATE  = "U"
	HEAD    = "H"
	OPTIONS = "O"
	DELETE  = "D"
)

func (router *Router) Get(path string, handler RoutingFunc) {
	router.AddBuiltinEndpoint(path, "GET", handler)
}

func (router *Router) Put(path string, handler RoutingFunc) {
	router.AddBuiltinEndpoint(path, "PUT", handler)
}

func (router *Router) Post(path string, handler RoutingFunc) {
	router.AddBuiltinEndpoint(path, "POST", handler)
}

func (router *Router) Update(path string, handler RoutingFunc) {
	router.AddBuiltinEndpoint(path, "UPDATE", handler)
}

func (router *Router) Head(path string, handler RoutingFunc) {
	router.AddBuiltinEndpoint(path, "HEAD", handler)
}

func (router *Router) Options(path string, handler RoutingFunc) {
	router.AddBuiltinEndpoint(path, "OPTIONS", handler)
}

func (router *Router) Delete(path string, handler RoutingFunc) {
	router.AddBuiltinEndpoint(path, "DELETE", handler)
}

func (router *Router) Any(path string, handler RoutingFunc) {
	router.AddBuiltinEndpoint(path, "*", handler)
}

func (router *Router) Namespace(name string) *Namespace {
	newNamespace := newNamespace(name, router.root)
	router.root.namespaces = append(router.root.namespaces, newNamespace)
	return newNamespace
}

func (router *Router) GetNamespace(name string) (nm *Namespace, found bool) {
	return router.root.GetNamespace(name)
}

func (router *Router) Root(handler RoutingFunc) {
	router.root.Root(handler)
}
