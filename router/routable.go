package router

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
