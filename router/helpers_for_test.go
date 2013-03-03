package router

import (
	"github.com/acsellers/thoreni"
	"net/http"
	"net/url"
)

func newTestContext(req *http.Request) *thoreni.Contextable {
	return &thoreni.Contextable{Renderable: new(testContext), Request: req}
}

type testContext struct {
	rendered string
}

func (t *testContext) Render(s string, d interface{}) {
	t.rendered = s
}
func (t *testContext) RenderStatic(s string, d interface{}) {
	t.rendered = s
}
func (t *testContext) Layout(s string) {
	t.rendered = s
}
func (t *testContext) Redirect(s string) {
	t.rendered = s
}

func requestish(path, method string) *http.Request {
	req := http.Request{Method: method}
	req.URL = &url.URL{Path: path}
	return &req
}
