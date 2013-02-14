package router

import "strings"

type MatchChecker interface {
	RespondsTo(Requestish) bool
}

type SimpleChecker struct {
	pattern string
	method  string
}

func (sc SimpleChecker) RespondsTo(rq Requestish) bool {
	return rq.Method() == sc.method && strings.HasPrefix(rq.Path(), sc.pattern)
}
