// Package router implements a fairly comprehensive URL router
package router

// Router is an HTTP request router, that takes existing net/http methods and paradigms and 
// extends them to make common actions easier to construct and give additional functionality
// to your Go code.
//
// It will redirect user requests to different functions,
// passing along the request and responsewriter in a library specific struct called
// Contextable. Contextable allows you to reuse existing code written against existing
// net/http contstructs, while giving you the flexibility to extend code to use new 
// functionality like the automatic template rendering and request format typing to render 
// different content based on what the request asks for. 
//
// Using the hooks library, you can 
// add in diffent bits of code to set variables, check cookie values to disallow access, 
// redirect requests to other paths, and other similar actions.
type Router struct {
	root      *Namespace
	generated bool
	//internalMap *RouteMap
	NotFound RoutingFunc
}

type Routable interface {
	Match(req Requestish) RoutingFunc
}

// NewRouter returns a new Router
func NewRouter() *Router {
	r := new(Router)
	r.root = newNamespace("", nil)

	r.NotFound = Default404

	return r
}

// Match will use the Router's tree of routes to find the correct route and returns the function 
// that corresponds to the correct route. The execution time will vary along with the number of
// routes defined.
func (router *Router) Match(req Requestish) RoutingFunc {
	matchingResponses := make([]*Endpoint, 0)
	if response, found := router.root.Serves(req); found {
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

func (router *Router) AddBuiltinEndpoint(path, method string, handler RoutingFunc) {
	if hasColonOperators(path) {
		regexChecker := NewRegexChecker(method, path)
		endpoint := &Endpoint{MatchChecker: regexChecker, RoutingFunc: handler, Name: path, rootedName: path}
		router.root.endpoints = append(router.root.endpoints, endpoint)
		return
	}
	checker := &SimpleChecker{pattern: path, method: method}
	endpoint := &Endpoint{MatchChecker: checker, RoutingFunc: handler, Name: path, rootedName: path}
	router.root.endpoints = append(router.root.endpoints, endpoint)
}
