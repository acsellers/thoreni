package router

type Endpoint struct {
	MatchChecker
	RoutingFunc
	Name       string
	rootedName string
}

func (ep *Endpoint) Serves(req Requestish) (found bool) {
	if ep.RespondsTo(req) {
		found = true
	}
	return
}
