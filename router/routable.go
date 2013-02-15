package router

func (router *Router) Get(path string, handler RoutingFunc) {
	getChecker := &SimpleChecker{pattern: path, method: "GET"}
	getEndpoint := &Endpoint{MatchChecker: getChecker, RoutingFunc: handler, Name: path, rootedName: path}
	router.Root.Endpoints = append(router.Root.Endpoints, getEndpoint)
}

func (router *Router) Put(path string, handler RoutingFunc) {
	putChecker := &SimpleChecker{pattern: path, method: "PUT"}
	putEndpoint := &Endpoint{MatchChecker: putChecker, RoutingFunc: handler, Name: path, rootedName: path}
	router.Root.Endpoints = append(router.Root.Endpoints, putEndpoint)
}

func (router *Router) Post(path string, handler RoutingFunc) {
	postChecker := &SimpleChecker{pattern: path, method: "POST"}
	postEndpoint := &Endpoint{MatchChecker: postChecker, RoutingFunc: handler, Name: path, rootedName: path}
	router.Root.Endpoints = append(router.Root.Endpoints, postEndpoint)
}

func (router *Router) Delete(path string, handler RoutingFunc) {
	deleteChecker := &SimpleChecker{pattern: path, method: "DELETE"}
	deleteEndpoint := &Endpoint{MatchChecker: deleteChecker, RoutingFunc: handler, Name: path, rootedName: path}
	router.Root.Endpoints = append(router.Root.Endpoints, deleteEndpoint)
}

func (router *Router) Any(path string, handler RoutingFunc) {
	deleteChecker := &SimpleChecker{pattern: path, method: "*"}
	deleteEndpoint := &Endpoint{MatchChecker: deleteChecker, RoutingFunc: handler, Name: path, rootedName: path}
	router.Root.Endpoints = append(router.Root.Endpoints, deleteEndpoint)
}
