package webserver

import "testing"

func TestResponse_StatusCode(t *testing.T) {
	response := response{
		statusCode: 200,
	}

	receivedStatusCode := response.StatusCode()
	if receivedStatusCode != 200 {
		t.Fatalf("Expected StatusCode() to return 200 but received %d", receivedStatusCode)
	}
}

func TestResponse_SetStatusCode(t *testing.T) {
	response := response{
		statusCode: 200,
	}

	response.SetStatusCode(404)
	if response.StatusCode() != 404 {
		t.Fatalf("Expected SetStatusCode() to set the status code to 404 but received %d", response.StatusCode())
	}
}

func TestResponse_Headers(t *testing.T) {
	response := response{
		headers: newResponseHeaders(),
	}

	receivedHeaders := response.Headers()
	if receivedHeaders == nil {
		t.Fatalf("Expected Headers() to return a non-nil value but received nil")
	}

	if receivedHeaders != response.headers {
		t.Fatalf("Expected Headers() to return the same value as the headers field but received a different value")
	}
}

func TestResponse_Body(t *testing.T) {
	response := response{
		body: []byte("Hello World!"),
	}

	receivedBody := response.Body()
	if string(receivedBody) != "Hello World!" {
		t.Fatalf("Expected Body() to return \"Hello World!\" but received \"%s\"", string(receivedBody))
	}
}

func TestResponse_SetBody(t *testing.T) {
	response := response{
		body: []byte("Hello World!"),
	}

	response.SetBody([]byte("Goodbye World!"))
	if string(response.Body()) != "Goodbye World!" {
		t.Fatalf("Expected SetBody() to set the body to \"Goodbye World!\" but received \"%s\"", string(response.Body()))
	}
}

func TestNewResponse(t *testing.T) {
	response := NewResponse(404)
	if response.StatusCode() != 404 {
		t.Fatalf("Expected NewResponse() to set the status code to 404 but received %d", response.StatusCode())
	}

	if response.Headers() == nil {
		t.Fatalf("Expected NewResponse() to set the headers to a non-nil value but received nil")
	}
}

func TestNewResponseWithBody(t *testing.T) {
	response := NewResponseWithBody(404, []byte("Hello World!"))
	if response.StatusCode() != 404 {
		t.Fatalf("Expected NewResponseWithBody() to set the status code to 404 but received %d", response.StatusCode())
	}

	if string(response.Body()) != "Hello World!" {
		t.Fatalf("Expected NewResponseWithBody() to set the body to \"Hello World!\" but received \"%s\"", string(response.Body()))
	}

	if response.Headers() == nil {
		t.Fatalf("Expected NewResponseWithBody() to set the headers to a non-nil value but received nil")
	}
}

func TestOkResponse(t *testing.T) {
	response := OkResponse()
	if response.StatusCode() != 200 {
		t.Fatalf("Expected OkResponse() to set the status code to 200 but received %d", response.StatusCode())
	}

	if response.Headers() == nil {
		t.Fatalf("Expected OkResponse() to set the headers to a non-nil value but received nil")
	}
}

func TestOkResponseWithBody(t *testing.T) {
	response := OkResponseWithBody([]byte("Hello World!"))
	if response.StatusCode() != 200 {
		t.Fatalf("Expected OkResponseWithBody() to set the status code to 200 but received %d", response.StatusCode())
	}

	if string(response.Body()) != "Hello World!" {
		t.Fatalf("Expected OkResponseWithBody() to set the body to \"Hello World!\" but received \"%s\"", string(response.Body()))
	}

	if response.Headers() == nil {
		t.Fatalf("Expected OkResponseWithBody() to set the headers to a non-nil value but received nil")
	}
}

func TestBadRequestResponse(t *testing.T) {
	response := BadRequestResponse()
	if response.StatusCode() != 400 {
		t.Fatalf("Expected BadRequestResponse() to set the status code to 400 but received %d", response.StatusCode())
	}

	if response.Headers() == nil {
		t.Fatalf("Expected BadRequestResponse() to set the headers to a non-nil value but received nil")
	}
}

func TestBadRequestResponseWithBody(t *testing.T) {
	response := BadRequestResponseWithBody([]byte("Hello World!"))
	if response.StatusCode() != 400 {
		t.Fatalf("Expected BadRequestResponseWithBody() to set the status code to 400 but received %d", response.StatusCode())
	}

	if string(response.Body()) != "Hello World!" {
		t.Fatalf("Expected BadRequestResponseWithBody() to set the body to \"Hello World!\" but received \"%s\"", string(response.Body()))
	}

	if response.Headers() == nil {
		t.Fatalf("Expected BadRequestResponseWithBody() to set the headers to a non-nil value but received nil")
	}
}

func TestNotFoundResponse(t *testing.T) {
	response := NotFoundResponse()
	if response.StatusCode() != 404 {
		t.Fatalf("Expected NotFoundResponse() to set the status code to 404 but received %d", response.StatusCode())
	}

	if response.Headers() == nil {
		t.Fatalf("Expected NotFoundResponse() to set the headers to a non-nil value but received nil")
	}
}

func TestNotFoundResponseWithBody(t *testing.T) {
	response := NotFoundResponseWithBody([]byte("Hello World!"))
	if response.StatusCode() != 404 {
		t.Fatalf("Expected NotFoundResponseWithBody() to set the status code to 404 but received %d", response.StatusCode())
	}

	if string(response.Body()) != "Hello World!" {
		t.Fatalf("Expected NotFoundResponseWithBody() to set the body to \"Hello World!\" but received \"%s\"", string(response.Body()))
	}

	if response.Headers() == nil {
		t.Fatalf("Expected NotFoundResponseWithBody() to set the headers to a non-nil value but received nil")
	}
}

func TestInternalErrorResponse(t *testing.T) {
	response := InternalErrorResponse()
	if response.StatusCode() != 500 {
		t.Fatalf("Expected InternalErrorResponse() to set the status code to 500 but received %d", response.StatusCode())
	}

	if response.Headers() == nil {
		t.Fatalf("Expected InternalErrorResponse() to set the headers to a non-nil value but received nil")
	}
}

func TestInternalErrorResponseWithBody(t *testing.T) {
	response := InternalErrorResponseWithBody([]byte("Hello World!"))
	if response.StatusCode() != 500 {
		t.Fatalf("Expected InternalErrorResponseWithBody() to set the status code to 500 but received %d", response.StatusCode())
	}

	if string(response.Body()) != "Hello World!" {
		t.Fatalf("Expected InternalErrorResponseWithBody() to set the body to \"Hello World!\" but received \"%s\"", string(response.Body()))
	}

	if response.Headers() == nil {
		t.Fatalf("Expected InternalErrorResponseWithBody() to set the headers to a non-nil value but received nil")
	}
}
