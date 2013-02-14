package router

import "testing"

func TestBasic(t *testing.T) {
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
