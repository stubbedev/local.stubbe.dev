package routes

import (
	"net/http"

	_template "github.com/stubbedev/local.stubbe.dev/internal/template"
)

func RouteHandler(route string) {
	routeHandler := func(wtr http.ResponseWriter, r *http.Request) {
		tpl := _template.LoadTemplate(route)
		_template.RenderTemplateData(wtr, tpl, nil)
	}
	http.HandleFunc(route, routeHandler)
}
