package main

import (
	_routes "github.com/stubbedev/local.stubbe.dev/internal/routes"
	_webserver "github.com/stubbedev/local.stubbe.dev/internal/webserver"
)

func main() {
	_routes.RouteHandler("/")
	_webserver.ServeEnv()
}
