package render

import (
	"bytes"
	"html/template"
	"path"
)

type TemplateRenderer struct {
	masterTemplate  *template.Template
	renderedStatics map[string]string
}

//TODO figure out if I need can reuse the same template after 
// calling ParseGlob or need to get the returned template
// also decide what do do about error from it
func NewRenderer(templatePath string) {
	r := new(TemplateRenderer)
	r.masterTemplate = template.New("master")
	r.masterTemplate.ParseGlob(path.Join(templatePath, templateGlob))
	r.renderedStatics = make(map[string]string)

	MasterRenderer = r
}

func NewRendererFromLists(layoutList, viewList []string) {
	r := new(TemplateRenderer)
	r.masterTemplate = template.New("master")
	for _, layout := range layoutList {
		r.masterTemplate.Parse(layout)
	}
	for _, view := range viewList {
		r.masterTemplate.Parse(view)
	}
	r.renderedStatics = make(map[string]string)

	MasterRenderer = r

}

func AddFolder(templatePath string) {
	MasterRenderer.masterTemplate.ParseGlob(path.Join(templatePath, templateGlob))
}

func (r TemplateRenderer) RenderTemplate(templateName, layoutName string, renderData interface{}) (header, page, footer string) {
	headerBuffer := new(bytes.Buffer)
	contentBuffer := new(bytes.Buffer)
	footerBuffer := new(bytes.Buffer)
	if MasterRenderer.masterTemplate.Lookup(layoutName+"-header") != nil {
		if err := MasterRenderer.masterTemplate.ExecuteTemplate(headerBuffer, layoutName+"-header", renderData); err == nil {
			header = string(headerBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}
	if MasterRenderer.masterTemplate.Lookup(layoutName+"-footer") != nil {
		if err := MasterRenderer.masterTemplate.ExecuteTemplate(footerBuffer, layoutName+"-footer", renderData); err == nil {
			footer = string(footerBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}
	if MasterRenderer.masterTemplate.Lookup(templateName) != nil {
		if err := MasterRenderer.masterTemplate.ExecuteTemplate(contentBuffer, templateName, renderData); err == nil {
			page = string(contentBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}
	return
}

func (r TemplateRenderer) RenderStatic(templateName, layoutName string, renderData interface{}) (header, content, footer string) {
	headerBuffer := new(bytes.Buffer)
	footerBuffer := new(bytes.Buffer)
	if MasterRenderer.masterTemplate.Lookup(layoutName+"-header") != nil {
		if err := MasterRenderer.masterTemplate.ExecuteTemplate(headerBuffer, layoutName+"-header", renderData); err == nil {
			header = string(headerBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}
	if MasterRenderer.masterTemplate.Lookup(layoutName+"-footer") != nil {
		if err := MasterRenderer.masterTemplate.ExecuteTemplate(footerBuffer, layoutName+"-footer", renderData); err == nil {
			footer = string(footerBuffer.Bytes())
		} else {
			//TODO log error or do something with it
		}
	}

	if precontent, found := MasterRenderer.renderedStatics[templateName]; found {
		content = precontent
	} else {
		contentBuffer := new(bytes.Buffer)
		if MasterRenderer.masterTemplate.Lookup(templateName) != nil {
			if err := MasterRenderer.masterTemplate.ExecuteTemplate(contentBuffer, templateName, renderData); err == nil {
				content = string(contentBuffer.Bytes())
			} else {
				//TODO log error or do something with it
			}
		}
	}
	return
}
