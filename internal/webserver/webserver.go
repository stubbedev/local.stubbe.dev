package webserver

import (
	"fmt"
	"log"
	"net/http"

	_os "github.com/stubbedev/local.stubbe.dev/internal/os"
)

func Serve(host string, port string) {
	host_and_port := host + ":" + port
	fmt.Printf("Listening on: %s", host_and_port)
	log.Fatal(
		http.ListenAndServe(host_and_port, nil),
	)
}

func ServeEnv() {
	port := _os.GetEnvVariable("PORT")
	host := _os.GetEnvVariable("HOST")
	Serve(host, port)
}
