package webserver

import (
	"bufio"
	"errors"
	"strings"
	"testing"
)

func TestHeaders_GetHeaderWithValidHeader(t *testing.T) {
	headers := headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	headerValue, err := headers.GetHeader("Host")
	if err != nil {
		t.Fatalf("Received error trying to get header %v", err)
	} else if headerValue != "www.bing.com" {
		t.Fatalf("Expected Host to have a value of www.bing.com but received %s", headerValue)
	}
}

func TestHeaders_GetHeaderWithValidDifferentCaseHeader(t *testing.T) {
	headers := headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	headerValue, err := headers.GetHeader("content-length")
	if err != nil {
		t.Fatalf("Received error trying to get header %v", err)
	} else if headerValue != "13" {
		t.Fatalf("Expected Host to have a value of 13 but received %s", headerValue)
	}
}

func TestHeaders_GetHeaderWithInvalidHeader(t *testing.T) {
	headers := headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	_, err := headers.GetHeader("Transfer-Encoding")
	if !errors.Is(err, ErrHeaderNotFound) {
		t.Fatalf("Expected the ErrHeaderNotFound error but received: %v", err)
	}
}

func TestHeaders_HasHeaderWithValidHeader(t *testing.T) {
	headers := headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	if !headers.HasHeader("Content-Length") {
		t.Fatalf("Could not find the Content-Length header")
	}
}

func TestHeaders_HasHeaderWithValidDifferentCaseHeader(t *testing.T) {
	headers := headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	if !headers.HasHeader("content-length") {
		t.Fatalf("Could not find the content-length header")
	}
}

func TestHeaders_HasHeaderWithInvalidHeader(t *testing.T) {
	headers := headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	if headers.HasHeader("Transfer-Encoding") {
		t.Fatalf("Found Transfer-Encoding header when it doesn't exist")
	}
}

func TestHeaders_GetAsMap(t *testing.T) {
	headers := &headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	headersMap := headers.GetAsMap()
	if len(headersMap) != 2 {
		t.Fatalf("Expected headers map to have 2 entries but had %d", len(headersMap))
	}

	if headersMap["Content-Length"] != "13" {
		t.Fatalf("Expected Content-Length to be 13 but was %s", headersMap["Content-Length"])
	}

	if headersMap["Host"] != "www.bing.com" {
		t.Fatalf("Expected Host to be www.bing.com but was %s", headersMap["Host"])
	}
}

func TestResponseHeaders_SetHeader(t *testing.T) {
	headers := &headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	headers.SetHeader("Content-Length", "15")
	if headers.headersMap["Content-Length"] != "15" {
		t.Fatalf("Expected Content-Length to be 15 but was %s", headers.headersMap["Content-Length"])
	}
}

func TestResponseHeaders_ClearHeaders(t *testing.T) {
	headers := &headers{
		headersMap: map[string]string{
			"Content-Length": "13",
			"Host":           "www.bing.com",
		},
	}

	headers.ClearHeaders()
	if len(headers.headersMap) != 0 {
		t.Fatalf("Expected headers to be empty but was %v", headers.headersMap)
	}
}

func TestParseHeaders(t *testing.T) {
	headersStream := strings.NewReader("Content-Length: 13\r\nHost: www.bing.com\r\n\r\n")

	headers, err := parseRequestHeaders(bufio.NewReader(headersStream))

	if err != nil {
		t.Fatalf("Received an error while parsing headers %v", err)
	}

	expectedContentLength := "13"
	parsedContentLength, err := headers.GetHeader("Content-Length")
	if err != nil {
		t.Fatalf("Received error while retrieving Content-Length header: %v", err)
	} else if expectedContentLength != parsedContentLength {
		t.Fatalf("Expected Content-Length to be %s but was %s", expectedContentLength, parsedContentLength)
	}

	expectedHost := "www.bing.com"
	parsedHost, err := headers.GetHeader("Host")
	if err != nil {
		t.Fatalf("Received error while retrieving Host header: %v", err)
	} else if expectedHost != parsedHost {
		t.Fatalf("Expected Host to be %s but was %s", expectedHost, parsedHost)
	}
}

func TestParseHeadersInvalidStream(t *testing.T) {
	headersStream := strings.NewReader("Content-Length: 13\r\nHost-www.bing.com\r\n\r\n")

	_, err := parseRequestHeaders(bufio.NewReader(headersStream))

	if !errors.Is(err, ErrInvalidHeader) {
		t.Fatalf("Expected an ErrInvalidHeader error but received %v", err)
	}
}
