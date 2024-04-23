package main

import (
	"fmt"
	"golang-webserver/webserver"
	"regexp"
)

func main() {
	ws := webserver.NewWebServer()

	// Add a handler with a `StringPath`
	ws.AddHandler(webserver.NewHandler(webserver.MethodGet, webserver.StringPath("/api/test"), func(path string) (int, []byte) {
		return 200, []byte("Hello World!")
	}))

	// Add a handler with a RegexPath
	ws.AddHandler(webserver.NewHandler(webserver.MethodAny, webserver.RegexPath(regexp.MustCompile("^/api.*$")), func(path string) (int, []byte) {
		return 200, []byte("Received request at " + path)
	}))

	// Add a static files handler
	ws.StaticFiles("www")

	if err := ws.Run(8080); err != nil {
		fmt.Printf("Fatal error occurred while running server: %v", err)
	}
}
