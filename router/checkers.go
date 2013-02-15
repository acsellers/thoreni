package router

import "regexp"
import "strings"

type MatchChecker interface {
	RespondsTo(Requestish) bool
	Surety(Requestish) int
}

type SimpleChecker struct {
	pattern string
	method  string
}

func (sc SimpleChecker) RespondsTo(rq Requestish) bool {
	return (rq.Method() == sc.method || sc.method == "*") && strings.HasPrefix(rq.Path(), sc.pattern)
}
func (sc SimpleChecker) Surety(rq Requestish) int {
	if sc.RespondsTo(rq) {
		return len(sc.pattern) + 1
	}
	if sc.method == "*" {
		return 1
	}
	return 2
}

type RegexChecker struct {
	simple *SimpleChecker
	regex  *regexp.Regexp
}

func NewRegexChecker(pattern, method, requestRegex string) (*RegexChecker, error) {
	regex, err := regexp.Compile(requestRegex)
	rc := &RegexChecker{simple: &SimpleChecker{pattern: pattern, method: method}, regex: regex}
	return rc, err
}

func (rc RegexChecker) RespondsTo(rq Requestish) bool {
	return true
}

//TODO: this is not taking into account the matched length, fix when we start dealing with the regex
func (rc RegexChecker) Surety(rq Requestish) int {
	return rc.simple.Surety(rq)
}
