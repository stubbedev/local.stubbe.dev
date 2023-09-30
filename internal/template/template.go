package template

import (
	"html/template"
	"net/http"
	"strings"
)

const PAGE_TEMPLATES = "static/pages"
const PAGE_PARTIALS = "static/partials"
const PAGE_ASSETS = "static/assets"

func LoadTemplate(route string) *template.Template {
	path := PAGE_TEMPLATES + route
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
