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

func TestBasicLayoutRenderSimple(t *testing.T) {
	mockRenderer := makeTestRenderer()
	mockRenderer.masterTemplate.Parse(`{{define "default-header"}}header{{end}}{{define "default-footer"}}footer{{end}}`)
	header, page, footer := mockRenderer.renderTemplate("test", "default", nil)
	if header != "header" {
		t.Fatal("Header for simple header should be '\"header\"'")
	}
	if page != "" {
		t.Fatal("Page for undefined template should be '\"\"'")
	}
	if footer != "footer" {
		t.Fatal("Footer for simple footershould be '\"footer\"'")
	}
}

func TestBasicViewRenderSimple(t *testing.T) {
	mockRenderer := makeTestRenderer()
	mockRenderer.masterTemplate.Parse(`{{define "default-header"}}header{{end}}{{define "default-footer"}}footer{{end}}`)
	mockRenderer.masterTemplate.Parse(`{{define "test"}}test{{end}}`)
	header, page, footer := mockRenderer.renderTemplate("test", "default", nil)
	if header != "header" {
		t.Fatal("Header for simple header should be '\"header\"'")
	}
	if page != "test" {
		t.Fatalf("Page for simple page should be '\"test\"' was '%s'", page)
	}
	if footer != "footer" {
		t.Fatal("Footer for simple footershould be '\"footer\"'")
	}
}
func makeTestRenderer() *templateRenderer {
	tr := new(templateRenderer)
	tr.masterTemplate = template.New("master")
	tr.renderedStatics = make(map[string]string)
	masterRenderer = tr
	return tr
}
