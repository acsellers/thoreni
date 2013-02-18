package router

import (
	"fmt"
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

func (nm *Namespace) Match(req Requestish) (response []*Endpoint) {
	for _, namespace := range nm.namespaces {
		if namespace.Contains(req) {
			if gottenResponse, ok := namespace.Serves(req); ok {
				response = append(response, gottenResponse...)
			}
		}
	}
	for _, endpoint := range nm.endpoints {
		if ok := endpoint.Serves(req); ok {
			response = append(response, endpoint)
		}
	}
	return
}
func (nm *Namespace) Contains(req Requestish) bool {
	return strings.HasPrefix(req.Path(), nm.rootedName)
}
func (nm *Namespace) Serves(req Requestish) (response []*Endpoint, found bool) {
	response = nm.Match(req)
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
