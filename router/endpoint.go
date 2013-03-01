package router

type Endpoint struct {
	MatchChecker
	RoutingFunc
	Name       string
	rootedName string
}
