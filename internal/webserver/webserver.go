package webserver

import (
	"fmt"
	"log"
	"net/http"

	_filesystem "github.com/stubbedev/local.stubbe.dev/internal/filesystem"
	_os "github.com/stubbedev/local.stubbe.dev/internal/os"
	_routes "github.com/stubbedev/local.stubbe.dev/internal/routes"
)

func Serve(host string, port string) {
	host_and_port := host + ":" + port
	fmt.Printf("Listening on: %s", host_and_port)
	log.Fatal(
		http.ListenAndServe(host_and_port, nil),
	)
}

func SetRoute(route string) any {
	_routes.RouteHandler(route)
	return nil
}

func SetStaticRoutes(root string) {
	file_path_routes := _filesystem.GetFolderPaths(root)
	for _, path := range file_path_routes {
		SetRoute(_filesystem.RemovePathPrefix(path, root))
	}
}

func ServeEnv() {
	port := _os.GetEnvVariable("PORT")
	host := _os.GetEnvVariable("HOST")
	Serve(host, port)
}
