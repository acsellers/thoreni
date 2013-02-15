// Package router implements a fairly comprehensive URL router
package router

// Router is the main interace to the functionality, you will draw your routes directly on the router
// or on a namespace that was drawn on the router
type Router struct {
	Root      *Namespace
	generated bool
	//internalMap *RouteMap
	NotFound RoutingFunc
}

func NewRouter() *Router {
	r := new(Router)
	r.Root = new(Namespace)
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
