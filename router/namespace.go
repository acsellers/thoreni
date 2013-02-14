package router

import (
	"fmt"
	"strings"
)

type Namespace struct {
	Endpoints  []*Endpoint
	Namespaces []*Namespace
	Name       string
	rootedName string
}

func NewNamespace(name string, parent *Namespace) (newNS *Namespace) {
	newNS = new(Namespace)
	newNS.Name = name
	if parent == nil {
		newNS.rootedName = "/" + name
	} else {
		newNS.rootedName = fmt.Sprintf("%s/%s", parent.rootedName, name)
	}

	return
}

func (nm *Namespace) Match(req Requestish) (response RoutingFunc) {
	for _, namespace := range nm.Namespaces {
		if namespace.Contains(req) {
			if gottenResponse, ok := namespace.Serves(req); ok {
				response = gottenResponse
				return
			}
		}
	}
	for _, endpoint := range nm.Endpoints {
		if gottenResponse, ok := endpoint.Serves(req); ok {
			response = gottenResponse
			return
		}
	}
	return nil
}
func (nm *Namespace) Contains(req Requestish) bool {
	return strings.HasPrefix(req.Path(), nm.rootedName)
}
func (nm *Namespace) Serves(req Requestish) (response RoutingFunc, found bool) {
	response = nm.Match(req)
	found = response != nil
	return
}
