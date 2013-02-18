package router

func (namespace *Namespace) Get(path string, handler RoutingFunc) {
	namespace.AddBuiltinEndpoint(path, "GET", handler)
}

func (namespace *Namespace) Put(path string, handler RoutingFunc) {
	namespace.AddBuiltinEndpoint(path, "PUT", handler)
}

func (namespace *Namespace) Post(path string, handler RoutingFunc) {
	namespace.AddBuiltinEndpoint(path, "POST", handler)
}

func (namespace *Namespace) Update(path string, handler RoutingFunc) {
	namespace.AddBuiltinEndpoint(path, "UPDATE", handler)
}

func (namespace *Namespace) Head(path string, handler RoutingFunc) {
	namespace.AddBuiltinEndpoint(path, "HEAD", handler)
}

func (namespace *Namespace) Options(path string, handler RoutingFunc) {
	namespace.AddBuiltinEndpoint(path, "OPTIONS", handler)
}

func (namespace *Namespace) Delete(path string, handler RoutingFunc) {
	namespace.AddBuiltinEndpoint(path, "DELETE", handler)
}

func (namespace *Namespace) Any(path string, handler RoutingFunc) {
	namespace.AddBuiltinEndpoint(path, "*", handler)
}

func (namespace *Namespace) Root(handler RoutingFunc) {
	rootChecker := &RootChecker{method: "*", normalizedName: namespace.rootedName}
	endpoint := &Endpoint{MatchChecker: rootChecker, RoutingFunc: handler}
	namespace.endpoints = append(namespace.endpoints, endpoint)
}
func (namespace *Namespace) Namespace(name string) *Namespace {
	newNamespace := newNamespace(name, namespace)
	namespace.namespaces = append(namespace.namespaces, newNamespace)
	return newNamespace
}

func (namespace *Namespace) GetNamespace(name string) (nm *Namespace, found bool) {
	found = true
	for _, internalNamespace := range namespace.namespaces {
		if internalNamespace.name == name {
			nm = internalNamespace
			return
		}
	}

	found = false
	return
}
