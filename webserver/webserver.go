package webserver

import (
	"fmt"
	"net"
)

type WebServer struct {
	handlers       []*Handler
	defaultHandler *Handler
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
			return fmt.Errorf("failed to accept request: %w", err)
		}

		go w.handle(conn)
	}
}

func (w *WebServer) AddHandler(handler *Handler) {
	w.handlers = append(w.handlers, handler)
}

func (w *WebServer) handle(conn net.Conn) {
	defer conn.Close()

	// By default, return an internal error if something goes wrong
	response := InternalErrorResponse()

	// Create a defer function that will write the response
	defer func() {
		err := writeResponse(conn, response)
		if err != nil {
			fmt.Printf("Error was returned while processing request: %v", err)
		}
	}()

	request, err := parseRequest(conn)
	if err != nil {
		fmt.Printf("Request could not be parsed: %v", err)
		return
	}

	// First look for an appropriate handler
	handlerFound := false
	var handler *Handler
	for _, h := range w.handlers {
		if h.Matches(request) {
			handler = h
			handlerFound = true
			break
		}
	}

	// If we don't find a specific handler, assign either a 404 or default handler
	if !handlerFound && w.defaultHandler == nil {
		handler = NewHandler(MethodAny, AnyPath(), func(request Request) Response {
			fmt.Printf("Handler could not be found for %v", request.Path())
			return NotFoundResponse()
		})
	} else if !handlerFound {
		handler = w.defaultHandler
	}

	// Execute the handler and assign the results
	response = handler.Execute(request)
}

func writeResponse(conn net.Conn, response Response) (finalErr error) {
	defer func() {
		// Handle any errors that might have occurred
		if r := recover(); r != nil {
			finalErr = fmt.Errorf("failed to write response: %v", r)
		}
	}()

	// Write the header
	_ = mustReturn(conn.Write([]byte(fmt.Sprintf("HTTP/1.1 %v\r\n", statusResponses[response.StatusCode()]))))

	// Loop through each of the headers and add them
	headersMap := response.Headers().GetAsMap()
	for k, v := range headersMap {
		_ = mustReturn(conn.Write([]byte(fmt.Sprintf("%v: %v\r\n", k, v))))
	}

	// Add the body
	_ = mustReturn(conn.Write([]byte("\r\n")))
	_ = mustReturn(conn.Write(response.Body()))

	return
}

func mustReturn[T interface{}](x T, err error) T {
	if err != nil {
		panic(err)
	}

	return x
}

func NewWebServer() WebServer {
	return WebServer{
		handlers:       make([]*Handler, 0, 10),
		defaultHandler: nil,
	}
}
