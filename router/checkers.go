package router

import "regexp"
import "strings"

// MatchChecker is the interface that is needed for the struct that is used to determine 
// whether the Route that the Router is looking at should be matched, and the match strength
// for that Route.
//
// Responds_to is permissive in how matches are counted. A "/" handler would match all routes
// so if there is a basic match, then it is not a problem to just respond with true, Surety
// is where you should determine how strong of a match it is. Matching actions should be fast,
// Surety actions can be slower.
//
// 
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
	strength := 1
	if sc.method == "*" {
		strength = 0
	}

	if sc.RespondsTo(rq) {
		return len(sc.pattern)*2 + strength
	}
	return 0
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
