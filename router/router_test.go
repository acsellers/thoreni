package router

import "testing"

func TestBasicGet(t *testing.T) {
	r := NewRouter()
	r.Get("/", func(ctx *Contextable) {
		ctx.Render("testing")
	})

	testReq := &testRequestish{path: "/", method: "GET"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "testing" {
			t.Fail()
		}
	}
}

func TestTwoGets(t *testing.T) {
	r := NewRouter()
	r.Get("/asdf", func(ctx *Contextable) {
		ctx.Render("testing")
	})
	r.Get("/", func(ctx *Contextable) {
		ctx.Render("error")
	})

	testReq := &testRequestish{path: "/asdf", method: "GET"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "testing" {
			t.Fail()
		}
	}
}

func TestTwoGetsReversed(t *testing.T) {
	r := NewRouter()
	r.Get("/", func(ctx *Contextable) {
		ctx.Render("error")
	})
	r.Get("/asdf", func(ctx *Contextable) {
		ctx.Render("testing")
	})

	testReq := &testRequestish{path: "/asdf", method: "GET"}
	tc := newTestContext(testReq)
	r.Match(testReq)(tc)
	if atc, ok := tc.Renderable.(*testContext); ok {
		if atc.rendered != "testing" {
			t.Fail()
		}
	}
}
func newTestContext(req Requestish) *Contextable {
	return &Contextable{Renderable: new(testContext), Requestish: req}
}

type testContext struct {
	rendered string
}

func (t *testContext) Render(s string) {
	t.rendered = s
}

func (t *testContext) Write(p []byte) (n int, err error) {
	return 0, nil
}

type testRequestish struct {
	path, method string
}

func (tr testRequestish) Path() string {
	return tr.path
}
func (tr testRequestish) Method() string {
	return tr.method
}
