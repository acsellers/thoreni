package router

type Router struct {
	Root      *Namespace
	generated bool
	//internalMap *RouteMap
	NotFound RoutingFunc
}

func NewRouter() *Router {
	r := new(Router)
	r.Root = new(Namespace) // appending to nil arrays works, plus empty string is what name/rootedName should be
	r.NotFound = Default404

	return r
}

func (router *Router) Match(req Requestish) RoutingFunc {
	if response, found := router.Root.Serves(req); found {
		return response
	}
	return router.NotFound
}
