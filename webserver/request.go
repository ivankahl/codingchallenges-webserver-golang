package webserver

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var ErrInvalidRequest = errors.New("the request is in the incorrect format")
var ErrInvalidBody = errors.New("body was invalid")
var ErrUnsupportedBody = errors.New("body format is not supported")

type Request interface {
	Path() string
	Method() Method
	Headers() RequestHeaders
	Body() []byte
	BodyAsString() string
}

type request struct {
	path    string
	method  Method
	headers Headers
	body    []byte
}

func (r *request) Path() string {
	return r.path
}

func (r *request) Method() Method {
	return r.method
}

func (r *request) Headers() RequestHeaders {
	return r.headers
}

func (r *request) Body() []byte {
	return r.body
}

func (r *request) BodyAsString() string {
	return string(r.body)
}

func parseRequest(requestStream io.Reader) (Request, error) {
	reader := bufio.NewReader(requestStream)

	// Read the first line with the method and path
	startLine, err := reader.ReadString('\n')
	if err != nil {
		return &request{}, ErrInvalidRequest
	}

	// Parse the first line
	startLineParts := strings.Split(strings.TrimSpace(startLine), " ")

	method, err := methodFromString(startLineParts[0])
	if err != nil {
		return &request{}, ErrInvalidMethod
	}

	path := startLineParts[1]

	headers, err := parseRequestHeaders(reader)
	if err != nil {
		return &request{}, err
	}

	body, err := retrieveBody(headers, reader)
	if err != nil {
		return &request{}, err
	}

	return &request{
		path:    path,
		method:  method,
		headers: headers,
		body:    body,
	}, nil
}

func retrieveBody(headers Headers, reader *bufio.Reader) ([]byte, error) {
	// First try parse chunked, i.e. where Transfer-Encoding: chunked
	transferEncoding, err := headers.GetHeader("Transfer-Encoding")
	if err == nil {
		if strings.ToLower(transferEncoding) == "chunked" {
			return retrieveWithChunkedEncoding(reader)
		}

		// We don't support any other Transfer-Encoding values
		return nil, ErrUnsupportedBody
	} else if !errors.Is(err, ErrHeaderNotFound) {
		return nil, err
	}

	// Then we try parse based on Content-Length
	_, err = headers.GetHeader("Content-Length")
	if err == nil {
		return retrieveWithContentLength(headers, reader)
	} else if !errors.Is(err, ErrHeaderNotFound) {
		return nil, err
	}

	// We assume no body was passed
	return nil, nil
}

func retrieveWithChunkedEncoding(reader *bufio.Reader) ([]byte, error) {
	body := make([]byte, 0)

	chunkLengthStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, ErrInvalidBody
	}
	chunkLengthStr = strings.TrimSpace(chunkLengthStr)

	for chunkLengthStr != "" && chunkLengthStr != "0" {
		chunkLength, err := strconv.Atoi(chunkLengthStr)
		if err != nil {
			return nil, ErrInvalidBody
		}

		// Read the next line
		nextBody := make([]byte, chunkLength)
		_, err = io.ReadFull(reader, nextBody)
		if err != nil {
			return nil, ErrInvalidBody
		}

		// Add the following line to the body
		body = append(body, nextBody...)

		// We read a blank line so we can get rid of the \r\n
		_, err = reader.ReadString('\n')
		if err != nil {
			return nil, ErrInvalidBody
		}

		// Read the next chunk length
		chunkLengthStr, err = reader.ReadString('\n')
		if err != nil {
			return nil, ErrInvalidBody
		}
		chunkLengthStr = strings.TrimSpace(chunkLengthStr)
	}

	return body, nil
}

func retrieveWithContentLength(headers Headers, reader *bufio.Reader) ([]byte, error) {
	contentLengthStr, err := headers.GetHeader("Content-Length")
	if err != nil {
		return nil, err
	}

	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return nil, ErrInvalidHeader
	}

	// Read the request body
	body := make([]byte, contentLength)
	_, err = io.ReadFull(reader, body)
	if err != nil {
		return nil, ErrInvalidBody
	}

	return body, nil
}
