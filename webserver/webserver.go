package webserver

import (
	"fmt"
	"net"
	"strings"
)

type WebServer struct {
	handlers                []Handler
	defaultHandlerSpecified bool
	defaultHandler          Handler
}

var statusResponses = map[int]string{
	200: "200 OK",
	202: "202 Accepted",
	400: "400 Bad Request",
	404: "404 Not Found",
	500: "500 Internal Server Error",
}

func (w *WebServer) StaticFiles(www string) {
	w.defaultHandler = NewStaticFileHandler(www)
	w.defaultHandlerSpecified = true
}

func (w *WebServer) Run(port int) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return fmt.Errorf("failed to start the web server: %w", err)
	}

	fmt.Printf("Running on port %v\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Errorf("failed to accept request: %w", err)
		}

		go w.handle(conn)
	}
}

func (w *WebServer) AddHandler(handler Handler) {
	w.handlers = append(w.handlers, handler)
}

func (w *WebServer) handle(conn net.Conn) {
	defer conn.Close()

	// By default, return an internal error if something goes wrong
	statusCode := 500
	headers := make(map[string]string)
	content := make([]byte, 0)

	// Create a defer function that will write the response
	defer func() {
		err := writeResponse(conn, statusCode, headers, content)
		if err != nil {
			fmt.Printf("Error was returned while processing request: %v", err)
		}
	}()

	// Parse the incoming request
	body := make([]byte, 1024)
	_, err := conn.Read(body)
	if err != nil {
		statusCode = 400
		return
	}
	bodyStr := string(body)

	// Parse the body string
	bodyParts := strings.Split(bodyStr, " ")
	method := bodyParts[0]
	requestPath := bodyParts[1]

	// First look for an appropriate handler
	handlerFound := false
	var handler Handler
	for _, h := range w.handlers {
		if h.Matches(method, requestPath) {
			handler = h
			handlerFound = true
			break
		}
	}

	// If we don't find a specific handler, assign either a 404 or default handler
	if !handlerFound && !w.defaultHandlerSpecified {
		handler = NewHandler(MethodAny, AnyPath(), func(path string) (int, []byte) {
			fmt.Printf("Handler could not be found for %v", path)
			return 404, nil
		})
	} else if !handlerFound {
		handler = w.defaultHandler
	}

	// Execute the handler and assign the results
	statusCode, content = handler.Execute(requestPath)
}

func writeResponse(conn net.Conn, statusCode int, headers map[string]string, content []byte) (finalErr error) {
	defer func() {
		// Handle any errors that might have occurred
		if r := recover(); r != nil {
			finalErr = fmt.Errorf("failed to write response: %v", r)
		}
	}()

	// Write the header
	_ = mustReturn(conn.Write([]byte(fmt.Sprintf("HTTP/1.1 %v\r\n", statusResponses[statusCode]))))

	// Loop through each of the headers and add them
	for k, v := range headers {
		_ = mustReturn(conn.Write([]byte(fmt.Sprintf("%v: %v\r\n", k, v))))
	}

	// Add the body
	_ = mustReturn(conn.Write([]byte("\r\n")))
	_ = mustReturn(conn.Write(content))

	return
}

func mustReturn[T interface{}](x T, err error) T {
	if err != nil {
		panic(err)
	}

	return x
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func NewWebServer() WebServer {
	return WebServer{
		handlers:                make([]Handler, 0, 10),
		defaultHandlerSpecified: false,
		defaultHandler:          Handler{},
	}
}
