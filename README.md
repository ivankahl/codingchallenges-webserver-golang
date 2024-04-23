# Coding Challenges: Build Your Own Web Server in Golang

This is my solution to the [Coding Challenges: Build Your Own Web Server](https://codingchallenges.fyi/challenges/challenge-webserver/) challenge. I've implemented a *very basic* HTTP Web Server in Golang using low level sockets.

## How It Works

To set up a new web server, create a new `WebServer` using the `NewWebServer` function:

```go
package main

import "golang-webserver/webserver"

func main() {
	ws := webserver.NewWebServer()
}
```

Once you've created a web server, you can attach handlers that will handle different requests. These handlers use Regex to match them with request paths. The code below shows how you can add handlers using the `RegexPath` and `StringPath` methods to specify the path:

```go
// Add a handler with a `StringPath` that handles `/api/test`
ws.AddHandler(webserver.NewHandler(webserver.MethodGet, webserver.StringPath("/api/test"), func(path string) (int, []byte) {
    return 200, []byte("Hello World!")
}))

// Add a handler with a RegexPath that handles `/api*`
ws.AddHandler(webserver.NewHandler(webserver.MethodAny, webserver.RegexPath(regexp.MustCompile("^/api.*$")), func(path string) (int, []byte) {
    return 200, []byte("Received request at " + path)
}))
```

You can also add a default `StaticFiles` handler that will map requests to files in a specified files:

```go
ws.StaticFiles("www")
```

The code above will map all requests that don't match other handlers to files in the "www" folder.

Finally, run the web server by calling `ws.Run` along with the desired port:

```go
if err := ws.Run(8080); err != nil {
    fmt.Printf("Fatal error occurred while running server: %v", err)
}
```

## Getting Started

If you want to run the project locally, clone this repository.

Once cloned, open a new terminal and run the program using the following `go` command:

```shell
go run .
```

You should see the following messsage appear:

```text
Running on port 8080
```
