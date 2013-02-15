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
	matchingResponses := make([]*Endpoint, 0)
	if response, found := router.Root.Serves(req); found {
		matchingResponses = append(matchingResponses, response...)
	}

	if len(matchingResponses) > 0 {
		var surestIndex, surestLength int
		for index, endpoint := range matchingResponses {
			if endpoint.Surety(req) > surestLength {
				surestLength = endpoint.Surety(req)
				surestIndex = index
			}
		}
		return matchingResponses[surestIndex].RoutingFunc
	}
	return router.NotFound
}
