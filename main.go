package main

import (
	"encoding/json"
	"fmt"
	"golang-webserver/webserver"
	"regexp"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	ws := webserver.NewWebServer()

	// Add a handler with a `StringPath`
	ws.AddHandler(webserver.NewHandler(webserver.MethodGet, webserver.StringPath("/api/person"), func(request webserver.Request) webserver.Response {
		person := Person{
			Name: "John Doe",
			Age:  30,
		}

		responseJson, err := json.Marshal(person)
		if err != nil {
			return webserver.InternalErrorResponse()
		}

		response := webserver.OkResponseWithBody(responseJson)
		response.Headers().SetHeader("Content-Type", "application/json")

		return response
	}))

	// Add a handler with a RegexPath
	ws.AddHandler(webserver.NewHandler(webserver.MethodAny, webserver.RegexPath(regexp.MustCompile("^/api.*$")), func(request webserver.Request) webserver.Response {
		return webserver.OkResponseWithBody([]byte("Received request at " + request.Path()))
	}))

	// Map static files
	ws.StaticFiles("www")

	if err := ws.Run(8080); err != nil {
		fmt.Printf("Fatal error occurred while running server: %v", err)
	}

}
