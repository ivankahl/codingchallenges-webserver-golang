package webserver

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type Method string

const (
	MethodGet    Method = "GET"
	MethodPost          = "POST"
	MethodPut           = "PUT"
	MethodPatch         = "PATCH"
	MethodDelete        = "DELETE"
	MethodAny           = "*"
)

type HandlerFunc func(path string) (int, []byte)

type Handler struct {
	method      Method
	pathPattern Path
	handler     HandlerFunc
}

func (h Handler) Matches(method string, path string) bool {
	return (method == string(h.method) || h.method == MethodAny) && h.pathPattern.Matches(path)
}

func (h Handler) Execute(path string) (int, []byte) {
	return h.handler(path)
}

func NewHandler(method Method, path Path, handler HandlerFunc) Handler {
	return Handler{
		method:      method,
		pathPattern: path,
		handler:     handler,
	}
}

func NewStaticFileHandler(wwwFilePath string) Handler {
	return NewHandler(MethodGet, AnyPath(), func(requestPath string) (int, []byte) {
		// First clean the path
		cleanedRequestPath := path.Clean(requestPath)

		// Account for `/`
		if cleanedRequestPath == "/" {
			cleanedRequestPath = "/index.html"
		}

		// Get the file path
		filePath := path.Join(wwwFilePath, cleanedRequestPath)

		// Check if the file exists and return a 404 if it doesn't
		_, err := os.Stat(filePath)
		if err != nil && errors.Is(err, os.ErrNotExist) {
			return 404, nil
		} else if err != nil {
			fmt.Printf("Internal error occurred while finding a static file: %v", err)
			return 500, nil
		}

		// Read the file
		fileContents, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Internal error occurred while reading a static file: %v", err)
			return 500, nil
		}

		return 200, fileContents
	})
}
