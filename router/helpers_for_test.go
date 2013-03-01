package router

func newTestContext(req Requestish) *Contextable {
	return &Contextable{Renderable: new(testContext), Requestish: req}
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

type testRequestish struct {
	path, method string
}

func (tr testRequestish) Path() string {
	return tr.path
}
func (tr testRequestish) Method() string {
	return tr.method
}
