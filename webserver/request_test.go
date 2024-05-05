package webserver

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRequest_GetPath(t *testing.T) {
	request := request{
		path: "/hello",
	}

	receivedPath := request.Path()
	if receivedPath != "/hello" {
		t.Fatalf("Expected GetPath() to return /hello but received %s", receivedPath)
	}
}

func TestRequest_GetMethod(t *testing.T) {
	request := request{
		method: MethodGet,
	}

	receivedMethod := request.Method()
	if receivedMethod != MethodGet {
		t.Fatalf("Expected GetMethod() to return %s but received %s", MethodGet, receivedMethod)
	}
}

func TestRequest_GetBody(t *testing.T) {
	request := request{
		body: []byte("Hello, World!"),
	}

	expectedBody := []byte("Hello, World!")
	receivedBody := request.Body()
	if !bytes.Equal(expectedBody, receivedBody) {
		t.Fatalf("Expected GetBody() to return %v but received %v", expectedBody, receivedBody)
	}
}

func TestRequest_GetBodyAsString(t *testing.T) {
	request := request{
		body: []byte("Hello, World!"),
	}

	expectedBody := "Hello, World!"
	receivedBody := request.BodyAsString()
	if expectedBody != receivedBody {
		t.Fatalf("Expected GetBodyAsString() to return %s but received %s", expectedBody, receivedBody)
	}
}

func TestParseRequestValid(t *testing.T) {
	requestStream := strings.NewReader("GET /hello HTTP/2\r\nHost: www.bing.com\r\nUser-Agent: curl/7.54.0\r\nContent-Length: 13\r\n\r\nHello, World!")

	parsedRequest, err := parseRequest(requestStream)
	if err != nil {
		t.Fatalf("Received the following error: %v", err)
	}

	var wantMethod Method = MethodGet
	if parsedRequest.Method() != wantMethod {
		t.Fatalf("Expected method is %v but received %v", wantMethod, parsedRequest.Method())
	}

	wantPath := "/hello"
	if parsedRequest.Path() != wantPath {
		t.Fatalf("Expected path is %v but received %v", wantPath, parsedRequest.Path())
	}

	if !parsedRequest.Headers().HasHeader("Host") || !parsedRequest.Headers().HasHeader("User-Agent") || !parsedRequest.Headers().HasHeader("Content-Length") {
		t.Fatalf("Expected the Host, User-Agent and Content-Length headers but didn't receive all of them")
	}

	wantBody := "Hello, World!"
	if parsedRequest.BodyAsString() != wantBody {
		t.Fatalf("Expected body is %s but received %s", wantBody, parsedRequest.BodyAsString())
	}
}

func TestParseRequestInvalidMethod(t *testing.T) {
	requestStream := strings.NewReader("DFLKS /hello HTTP/2\r\nHost: www.bing.com\r\nUser-Agent: curl/7.54.0\r\nContent-Length: 13\r\n\r\nHello, World!")

	_, err := parseRequest(requestStream)

	if !errors.Is(err, ErrInvalidMethod) {
		t.Fatalf("Expected an ErrInvalidMethod error but received %v", err)
	}
}

func TestParseRequestInvalidHeader(t *testing.T) {
	requestStream := strings.NewReader("GET /hello HTTP/2\r\nHost: www.bing.com\r\nUser-Agent-curl/7.54.0\r\nContent-Length: 13\r\n\r\nHello, World!")

	_, err := parseRequest(requestStream)

	if !errors.Is(err, ErrInvalidHeader) {
		t.Fatalf("Expected an ErrInvalidHeader error but received %v", err)
	}
}

func TestRetrieveWithChunkedEncoding(t *testing.T) {
	bodyStream := strings.NewReader("5\r\nHello\r\n11\r\n there Ivan\r\n0\r\n")

	body, err := retrieveWithChunkedEncoding(bufio.NewReader(bodyStream))

	if err != nil {
		t.Fatalf("Received an error while retrieving chunked encoding when none was expected: %v", err)
	}

	if string(body) != "Hello there Ivan" {
		t.Fatalf("Expected \"Hello there Ivan\" but received \"%s\"", string(body))
	}
}

func TestRetrieveWithChunkedEncodingInvalidBody(t *testing.T) {
	bodyStream := strings.NewReader("Hello world\r\n")

	_, err := retrieveWithChunkedEncoding(bufio.NewReader(bodyStream))

	if !errors.Is(err, ErrInvalidBody) {
		t.Fatalf("Expected an ErrInvalidBody but received %v", err)
	}
}

func TestRetrieveWithContentLength(t *testing.T) {
	bodyStream := strings.NewReader("Hello World!")
	headers := &headers{
		headersMap: map[string]string{"Content-Length": "12"},
	}

	body, err := retrieveWithContentLength(headers, bufio.NewReader(bodyStream))

	if err != nil {
		t.Fatalf("Received an error while retrieving by content-length when none was expected %v", err)
	}

	if string(body) != "Hello World!" {
		t.Fatalf("Expected \"Hello World!\" but received \"%s\"", string(body))
	}
}

func TestRetrieveWithContentLengthInvalidBody(t *testing.T) {
	bodyStream := strings.NewReader("Hello!")
	headers := &headers{
		headersMap: map[string]string{"Content-Length": "12"},
	}

	_, err := retrieveWithContentLength(headers, bufio.NewReader(bodyStream))

	if !errors.Is(err, ErrInvalidBody) {
		t.Fatalf("Expected an ErrInvalidBody but received %v", err)
	}
}

func TestRetrieveWithContentLengthInvalidHeader(t *testing.T) {
	bodyStream := strings.NewReader("Hello!")
	headers := &headers{
		headersMap: map[string]string{"Content-Length": "Hi"},
	}

	_, err := retrieveWithContentLength(headers, bufio.NewReader(bodyStream))

	if !errors.Is(err, ErrInvalidHeader) {
		t.Fatalf("Expected an ErrInvalidHeader but received %v", err)
	}
}
