package router

type Endpoint struct {
	MatchChecker
	RoutingFunc
	Name       string
	rootedName string
}

func (ep *Endpoint) Serves(req Requestish) (response RoutingFunc, found bool) {
	if ep.RespondsTo(req) {
		response = ep.RoutingFunc
		found = true
	}
	return
}
