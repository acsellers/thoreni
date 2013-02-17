package router

import "fmt"
import "regexp"
import "strings"

var colonChecker = regexp.MustCompile(":id|:key|:slug|:name")

// MatchChecker is the interface that is needed for the struct that is used to determine 
// whether the Route that the Router is looking at should be matched, and the match strength
// for that Route.
//
// Responds_to is permissive in how matches are counted. A "/" handler would match all routes
// so if there is a basic match, then it is not a problem to just respond with true, Surety
// is where you should determine how strong of a match it is. Matching actions should be fast,
// Surety actions can be slower. Exact method matchers will rank 1 higher than matches on the any 
// method, we do this by doubling the length of the match.
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
	method string
	regex  *regexp.Regexp
}

func NewRegexChecker(method, requestRegex string) *RegexChecker {
	regex := regexp.MustCompile(ReplaceColonOperators(requestRegex))
	rc := &RegexChecker{method: method, regex: regex}
	return rc
}

func (rc RegexChecker) RespondsTo(rq Requestish) bool {
	return (rc.method == rq.Method() || rc.method == "*") && rc.regex.MatchString(rq.Path())
}

//TODO: this is not taking into account the matched length, fix when we start dealing with the regex
func (rc RegexChecker) Surety(rq Requestish) int {
	return len(rc.regex.FindString(rq.Path()))*2 + 1
}

func ReplaceColonOperators(pre string) (post string) {
	post = strings.Replace(pre, ":id", "([0-9]+)", -1)
	post = strings.Replace(post, ":key", "([0-9a-fA-F]+)", -1)
	post = strings.Replace(post, ":slug", "([a-fA-F0-9_-]+)", -1)
	post = strings.Replace(post, ":name", "([^/]+)", -1)
	post = fmt.Sprintf("^%s", post)
	return
}

func hasColonOperators(path string) bool {
	return colonChecker.MatchString(path)
}
