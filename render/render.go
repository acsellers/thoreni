package render

import "fmt"

var (
	MasterRenderer *TemplateRenderer
	templateGlob   string
)

func init() {
	SetFileType("tmpl")
}

func SetFileType(ending string) {
	templateGlob = fmt.Sprintf("*.%s", ending)
}
