package render

import (
	"html/template"
	"testing"
)

func TestEmptyRenderSimple(t *testing.T) {
	mockRenderer := makeTestRenderer()
	header, page, footer := mockRenderer.renderTemplate("test", "default", nil)
	if header != "" {
		t.Fatal("Header for empty template should be '\"\"'")
	}
	if page != "" {
		t.Fatal("Page for empty template should be '\"\"'")
	}
	if footer != "" {
		t.Fatal("Footer for empty template should be '\"\"'")
	}
}

func makeTestRenderer() *templateRenderer {
	tr := new(templateRenderer)
	tr.masterTemplate = template.New("master")
	tr.renderedStatics = make(map[string]string)
	masterRenderer = tr
	return tr
}
