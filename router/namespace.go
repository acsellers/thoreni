package router

import (
	"fmt"
	"github.com/acsellers/thoreni"
	"strings"
)

type Namespace struct {
	endpoints  []*Endpoint
	namespaces []*Namespace
	name       string
	rootedName string
	root       *Endpoint
}

func newNamespace(name string, parent *Namespace) (newNS *Namespace) {
	newNS = new(Namespace)
	newNS.name = name
	if parent == nil {
		newNS.rootedName = "/" + name
	} else {
		newNS.rootedName = fmt.Sprintf("%s/%s", parent.rootedName, name)
	}

	return
}

func (namespace *Namespace) Match(req thoreni.Requestish) (response []*Endpoint) {
	for _, namespace := range namespace.namespaces {
		if namespace.Contains(req) {
			if gottenResponse, ok := namespace.RespondsTo(req); ok {
				response = append(response, gottenResponse...)
			}
		}
	}
	for _, endpoint := range namespace.endpoints {
		if ok := endpoint.RespondsTo(req); ok {
			response = append(response, endpoint)
		}
	}
	return
}
func (namespace *Namespace) Contains(req thoreni.Requestish) bool {
	return strings.HasPrefix(req.Path(), namespace.rootedName)
}
func (namespace *Namespace) RespondsTo(req thoreni.Requestish) (response []*Endpoint, found bool) {
	response = namespace.Match(req)
	found = len(response) > 0
	return
}

func (namespace *Namespace) AddBuiltinEndpoint(path, method string, handler RoutingFunc) {
	if hasColonOperators(path) {
		regexChecker := NewRegexChecker(method, path)
		endpoint := &Endpoint{MatchChecker: regexChecker, RoutingFunc: handler, Name: path, rootedName: namespace.rootedName + path}
		namespace.endpoints = append(namespace.endpoints, endpoint)
		return
	}
	checker := &SimpleChecker{pattern: path, method: method}
	endpoint := &Endpoint{MatchChecker: checker, RoutingFunc: handler, Name: path, rootedName: namespace.rootedName + path}
	namespace.endpoints = append(namespace.endpoints, endpoint)
}
