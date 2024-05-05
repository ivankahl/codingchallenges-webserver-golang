package webserver

import (
	"bufio"
	"errors"
	"strings"
)

// ErrHeaderNotFound is returned when a header is not found in the headers map
var ErrHeaderNotFound = errors.New("the specified header could not be found")

// ErrInvalidHeader is returned when a header is not in the correct format
var ErrInvalidHeader = errors.New("the header was not in the correct format")

// Headers defines the methods available on a headers object. This is implemented by a local struct.
type Headers interface {
	// GetHeader returns the value of the specified header. The header key is case-insensitive
	GetHeader(header string) (string, error)
	// HasHeader returns whether the specified header exists. The header key is case-insensitive
	HasHeader(header string) bool
	// GetAsMap returns the headers as a map
	GetAsMap() map[string]string
}

// ResponseHeaders defines the methods available on a response headers object. This is implemented by a local struct.
type ResponseHeaders interface {
	// Headers is an extension of the Headers interface
	Headers
	// SetHeader sets the value of the specified header
	SetHeader(header string, value string)
	// ClearHeaders clears all headers
	ClearHeaders()
}

// RequestHeaders defines the methods available on a request headers object. This is implemented by a local struct.
type RequestHeaders interface {
	// Headers is an extension of the Headers interface
	Headers
}

// headers is a local struct that implements the Headers interface. It contains a map of headers.
type headers struct {
	// The map of headers
	headersMap map[string]string
}

// GetHeader returns the value of the specified header. The header key is case-insensitive
func (h *headers) GetHeader(header string) (string, error) {
	for k, v := range h.headersMap {
		if strings.ToLower(k) == strings.ToLower(header) {
			return v, nil
		}
	}

	return "", ErrHeaderNotFound
}

// HasHeader returns whether the specified header exists. The header key is case-insensitive
func (h *headers) HasHeader(header string) bool {
	for k := range h.headersMap {
		if strings.ToLower(k) == strings.ToLower(header) {
			return true
		}
	}

	return false
}

// GetAsMap returns the headers as a map
func (h *headers) GetAsMap() map[string]string {
	toReturn := make(map[string]string)

	for k, v := range h.headersMap {
		toReturn[k] = v
	}

	return toReturn
}

// SetHeader sets the value of the specified header
func (h *headers) SetHeader(header string, value string) {
	h.headersMap[header] = value
}

// ClearHeaders clears all headers
func (h *headers) ClearHeaders() {
	h.headersMap = make(map[string]string)
}

// newResponseHeaders creates a new response headers object
func newResponseHeaders() ResponseHeaders {
	return &headers{
		headersMap: make(map[string]string),
	}
}

// newRequestHeaders creates a new request headers object from an incoming HTTP request stream
func parseRequestHeaders(reader *bufio.Reader) (RequestHeaders, error) {
	headersMap := make(map[string]string)

	for {
		rawHeader, err := reader.ReadString('\n')
		if err != nil {
			return nil, ErrInvalidHeader
		}

		if strings.TrimSpace(rawHeader) == "" {
			break
		}

		rawHeaderParts := strings.Split(strings.TrimSpace(rawHeader), ": ")
		if len(rawHeaderParts) != 2 {
			return nil, ErrInvalidHeader
		}

		headersMap[rawHeaderParts[0]] = strings.TrimSpace(rawHeaderParts[1])
	}

	return &headers{
		headersMap: headersMap,
	}, nil
}
