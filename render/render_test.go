package render

import (
	"bytes"
	"html/template"
	"net/http"
	"testing"
)

func TestEmptyRenderSimple(t *testing.T) {
	mockRenderer := makeTestRenderer()
	header, page, footer := mockRenderer.RenderTemplate("test", "default", nil)
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
	header, page, footer := mockRenderer.RenderTemplate("test", "default", nil)
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
	header, page, footer := mockRenderer.RenderTemplate("test", "default", nil)
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

func TestBasicViewRenderStatic(t *testing.T) {
	mockRenderer := makeTestRenderer()
	mockRenderer.masterTemplate.Parse(`{{define "default-header"}}header{{end}}{{define "default-footer"}}footer{{end}}`)
	mockRenderer.renderedStatics["test"] = "test"
	header, page, footer := mockRenderer.RenderStatic("test", "default", nil)
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
func TestRenderStaticError(t *testing.T) {
	mockRenderer := makeTestRenderer()
	mockRenderer.masterTemplate.Parse(`{{define "default-header"}}header{{end}}{{define "default-footer"}}footer{{end}}`)
	rr := new(RequestRenderer)
	rr.master = mockRenderer
	rr.layout = "default"
	mrw := newMockResponseWriter()
	rr.Output = mrw
	rr.RenderStatic("test")
	if string(mrw.buffer.Bytes()) != "Template render error" {
		t.Fatalf("Exepecting 'Template Render Error' , was '%s'", string(mrw.buffer.Bytes()))
	}
}
func makeTestRenderer() *TemplateRenderer {
	tr := new(TemplateRenderer)
	tr.masterTemplate = template.New("master")
	tr.renderedStatics = make(map[string]string)
	MasterRenderer = tr
	return tr
}

type mockResponseWriter struct {
	buffer *bytes.Buffer
	header int
}

func (m mockResponseWriter) Write(b []byte) (int, error) {
	return m.buffer.Write(b)
}
func (m mockResponseWriter) Header() http.Header {
	return http.Header{}
}
func (m *mockResponseWriter) WriteHeader(status int) {
	m.header = status
}

func newMockResponseWriter() *mockResponseWriter {
	mrw := new(mockResponseWriter)
	mrw.buffer = new(bytes.Buffer)
	return mrw
}
