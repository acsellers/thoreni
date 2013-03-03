package render

import (
	"fmt"
	"net/http"
)

type RequestRenderer struct {
	master     *TemplateRenderer
	layout     string
	Output     http.ResponseWriter
	Request    *http.Request
	RenderData interface{}
}

func NewRequestRenderer() *RequestRenderer {
	rr := new(RequestRenderer)
	rr.master = MasterRenderer
	rr.layout = "default"
	return rr
}

func (r RequestRenderer) Render(templateName string) {
	if r.Output != nil {
		top, content, bottom := MasterRenderer.RenderTemplate(templateName, r.layout, r.RenderData)
		if top == "" || bottom == "" {
			//TODO write some kind of 500 error
			fmt.Fprint(r.Output, "Template render error")
		} else {
			if content == "" {
				//TODO write some kind of sorta 500 error
				fmt.Fprint(r.Output, "Content render error")
			} else {
				fmt.Fprint(r.Output, top, content, bottom)
			}
		}
	}
	return
}

func (r RequestRenderer) RenderStatic(templateName string) {
	if r.Output != nil {
		top, content, bottom := MasterRenderer.RenderStatic(templateName, r.layout, r.RenderData)
		if top == "" || bottom == "" {
			//TODO write some kind of 500 error
			fmt.Fprint(r.Output, "Template render error")
		} else {
			if content == "" {
				//TODO write some kind of sorta 500 error
				fmt.Fprint(r.Output, "Content render error")
			} else {
				fmt.Fprint(r.Output, top, content, bottom)
			}
		}
	}
}

func (r *RequestRenderer) Layout(layoutName string) {
	r.layout = layoutName
}

func (r *RequestRenderer) Redirect(address string) {
	http.Redirect(r.Output, r.Request, address, http.StatusSeeOther)
}
