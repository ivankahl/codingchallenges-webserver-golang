package webserver

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type HandlerFunc func(request Request) Response

type Handler struct {
	method      Method
	pathPattern Path
	handler     HandlerFunc
}

func (h *Handler) Matches(request Request) bool {
	return (request.Method() == h.method || h.method == MethodAny) && h.pathPattern.Matches(request.Path())
}

func (h *Handler) Execute(request Request) Response {
	return h.handler(request)
}

func NewHandler(method Method, path Path, handler HandlerFunc) *Handler {
	return &Handler{
		method:      method,
		pathPattern: path,
		handler:     handler,
	}
}

func NewStaticFileHandler(wwwFilePath string) *Handler {
	return NewHandler(MethodGet, AnyPath(), func(request Request) Response {
		// First clean the path
		cleanedRequestPath := path.Clean(request.Path())

		// Account for `/`
		if cleanedRequestPath == "/" {
			cleanedRequestPath = "/index.html"
		}

		// Get the file path
		filePath := path.Join(wwwFilePath, cleanedRequestPath)

		// Check if the file exists and return a 404 if it doesn't
		_, err := os.Stat(filePath)
		if err != nil && errors.Is(err, os.ErrNotExist) {
			return NotFoundResponse()
		} else if err != nil {
			fmt.Printf("Internal error occurred while finding a static file: %v", err)
			return InternalErrorResponse()
		}

		// Read the file
		fileContents, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Internal error occurred while reading a static file: %v", err)
			return InternalErrorResponse()
		}

		return OkResponseWithBody(fileContents)
	})
}
