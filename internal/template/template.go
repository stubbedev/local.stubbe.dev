package template

import (
	"html/template"
	"net/http"
	"strings"
)

func LoadTemplate(route string) *template.Template {
	path := "static/templates" + route
	var abs_path string
	if strings.HasSuffix(path, "/") {
		abs_path = path + "index.html"
	} else {
		abs_path = path + "/index.html"
	}
	tpl := template.Must(template.ParseFiles(abs_path))
	return tpl
}

func RenderTemplateData(wtr http.ResponseWriter, tpl *template.Template, inf any) any {
	render := tpl.Execute(wtr, inf)
	return render
}
