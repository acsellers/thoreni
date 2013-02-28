package render

import "bytes"
import "fmt"
import "html/template"
import "io"
import "path"

var (
	masterRenderer *templateRenderer
	templateGlob   string
)

func init() {
	SetFileType("tmpl")
}

func SetFileType(ending string) {
	templateGlob = fmt.Sprintf("*.%s", ending)
}

type templateRenderer struct {
	masterTemplate  *template.Template
	renderedStatics map[string]string
}
type RequestRenderer struct {
	master     *templateRenderer
	layout     string
	Output     io.Writer
	RenderData interface{}
}

//TODO figure out if I need can reuse the same template after 
// calling ParseGlob or need to get the returned template
// also decide what do do about error from it
func NewRenderer(templatePath string) {
	r := new(templateRenderer)
	r.masterTemplate = template.New("master")
	r.masterTemplate.ParseGlob(path.Join(templatePath, templateGlob))
	r.renderedStatics = make(map[string]string)

	masterRenderer = r
}

func NewRequestRenderer() *RequestRenderer {
	rr := new(RequestRenderer)
	rr.master = masterRenderer
	rr.layout = "default"
	return rr
}

func AddFolder(templatePath string) {
	masterRenderer.masterTemplate.ParseGlob(path.Join(templatePath, templateGlob))
}

func (r templateRenderer) renderTemplate(templateName, layoutName string, renderData interface{}) (header, page, footer string) {
	headerBuffer := new(bytes.Buffer)
	contentBuffer := new(bytes.Buffer)
	footerBuffer := new(bytes.Buffer)
	if masterRenderer.masterTemplate.Lookup(layoutName+"-header") != nil {
		if err := masterRenderer.masterTemplate.ExecuteTemplate(headerBuffer, layoutName+"-header", renderData); err == nil {
			header = string(headerBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}
	if masterRenderer.masterTemplate.Lookup(layoutName+"-footer") != nil {
		if err := masterRenderer.masterTemplate.ExecuteTemplate(footerBuffer, layoutName+"-footer", renderData); err == nil {
			footer = string(footerBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}
	if masterRenderer.masterTemplate.Lookup(templateName) != nil {
		if err := masterRenderer.masterTemplate.ExecuteTemplate(contentBuffer, templateName, renderData); err == nil {
			page = string(contentBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}
	return
}

func (r templateRenderer) renderStatic(templateName, layoutName string, renderData interface{}) (header, content, footer string) {
	headerBuffer := new(bytes.Buffer)
	footerBuffer := new(bytes.Buffer)
	if masterRenderer.masterTemplate.Lookup(layoutName+"-header") != nil {
		if err := masterRenderer.masterTemplate.ExecuteTemplate(headerBuffer, layoutName+"-header", renderData); err == nil {
			header = string(headerBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}
	if masterRenderer.masterTemplate.Lookup(layoutName+"-footer") != nil {
		if err := masterRenderer.masterTemplate.ExecuteTemplate(footerBuffer, layoutName+"-footer", renderData); err == nil {
			footer = string(footerBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}

	if precontent, found := masterRenderer.renderedStatics[templateName]; found {
		content = precontent
	} else {
		contentBuffer := new(bytes.Buffer)
		if masterRenderer.masterTemplate.Lookup(templateName) != nil {
			if err := masterRenderer.masterTemplate.ExecuteTemplate(contentBuffer, templateName, renderData); err == nil {
				content = string(contentBuffer.Bytes())
			} else {
				//TODO log error or do something with it
			}
		}
	}
	return
}

func (r RequestRenderer) Render(templateName string) {
	if r.Output != nil {
		top, content, bottom := masterRenderer.renderTemplate(templateName, r.layout, r.RenderData)
		if top == "" || bottom == "" {
			//TODO write some kind of 500 error
			fmt.Fprint(r.Output, "Template render error")
		} else {
			if content == "" {
				//TODO write some kind of sorta 500 error
				fmt.Fprint(r.Output, "Template render error")
			} else {
				fmt.Fprint(r.Output, top, content, bottom)
			}
		}
	}
	return
}

func (r RequestRenderer) RenderStatic(templateName string) {
	if r.Output != nil {
		top, content, bottom := masterRenderer.renderStatic(templateName, r.layout, r.RenderData)
		if top == "" || bottom == "" {
			//TODO write some kind of 500 error
			fmt.Fprint(r.Output, "Template render error")
		} else {
			if content == "" {
				//TODO write some kind of sorta 500 error
				fmt.Fprint(r.Output, "Template render error")
			} else {
				fmt.Fprint(r.Output, top, content, bottom)
			}
		}
	}
}

func (r *RequestRenderer) Layout(layoutName string) {
	r.layout = layoutName
}
