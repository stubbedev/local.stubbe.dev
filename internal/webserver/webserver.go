package webserver

import (
	"fmt"
	"log"
	"net/http"

	_filesystem "github.com/stubbedev/local.stubbe.dev/internal/filesystem"
	_os "github.com/stubbedev/local.stubbe.dev/internal/os"
	_routes "github.com/stubbedev/local.stubbe.dev/internal/routes"
	_template "github.com/stubbedev/local.stubbe.dev/internal/template"
)

func Serve(host string, port string) {
	SetStaticAssetsRoute()
	host_and_port := host + ":" + port
	fmt.Printf("Listening on: %s", host_and_port)
	log.Fatal(
		http.ListenAndServe(host_and_port, nil),
	)
}

func SetStaticAssetsRoute() {
	fs := http.FileServer(http.Dir("/static/assets/css"))
	// asset_paths := "/" + _template.PAGE_ASSETS + "/"
	http.Handle("/static/css", fs)
}

func SetRoute(route string) {
	_routes.RouteHandler(route)
}

func SetStaticRoutes() {
	root := _template.PAGE_TEMPLATES
	file_path_routes := append(_filesystem.GetFolderPaths(root), "/")
	for _, path := range file_path_routes {
		SetRoute(_filesystem.RemovePathPrefix(path, root))
	}
}

func ServeEnv() {
	port := _os.GetEnvVariable("PORT")
	host := _os.GetEnvVariable("HOST")
	Serve(host, port)
}
