package main

import (
	_webserver "github.com/stubbedev/local.stubbe.dev/internal/webserver"
)

func main() {
	_webserver.SetStaticRoutes()
	_webserver.ServeEnv()
}
