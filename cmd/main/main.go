package main

import (
	_webserver "github.com/stubbedev/local.stubbe.dev/internal/webserver"
)

func main() {
	_webserver.SetStaticRoutes("static/templates/")
	_webserver.ServeEnv()
}
