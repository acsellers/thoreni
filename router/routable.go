package router

func (router *Router) Get(path string, handler RoutingFunc) {
	getChecker := &SimpleChecker{pattern: path, method: "GET"}
	getEndpoint := &Endpoint{MatchChecker: getChecker, RoutingFunc: handler, Name: path, rootedName: path}
	router.Root.Endpoints = append(router.Root.Endpoints, getEndpoint)
}
