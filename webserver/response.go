package webserver

// Response defines the methods available on a response object. This is implemented by a local struct.
// A collection of methods are available to set the status code, headers, and body of the response when creating it.
type Response interface {
	// StatusCode returns the status code of the response
	StatusCode() int
	// SetStatusCode sets the status code of the response
	SetStatusCode(statusCode int)
	// Headers returns the headers of the response. This is a pointer so you can modify the headers directly.
	Headers() ResponseHeaders
	// Body returns the content of the response
	Body() []byte
	// SetBody sets the content of the response
	SetBody(content []byte)
}

// response is a local struct that implements the Response interface. It contains fields for the status code, headers,
// and content of the response.
type response struct {
	// The status code of the response
	statusCode int
	// The headers of the response
	headers ResponseHeaders
	// The content of the response
	body []byte
}

// StatusCode returns the status code of the response
func (r *response) StatusCode() int {
	return r.statusCode
}

// SetStatusCode sets the status code of the response
func (r *response) SetStatusCode(statusCode int) {
	r.statusCode = statusCode
}

// Headers returns the headers of the response. This is a pointer so you can modify the headers directly.
func (r *response) Headers() ResponseHeaders {
	return r.headers
}

// Body returns the content of the response
func (r *response) Body() []byte {
	return r.body
}

// SetBody sets the content of the response
func (r *response) SetBody(content []byte) {
	r.body = content
}

// NewResponse creates a new response with the given status code and no body
func NewResponse(statusCode int) Response {
	return &response{
		statusCode: statusCode,
		headers:    newResponseHeaders(),
	}
}

// NewResponseWithBody creates a new response with the given status code and body
func NewResponseWithBody(statusCode int, body []byte) Response {
	return &response{
		statusCode: statusCode,
		body:       body,
		headers:    newResponseHeaders(),
	}
}

// OkResponse creates a new response with a status code of 200
func OkResponse() Response {
	return NewResponse(200)
}

// OkResponseWithBody creates a new response with a status code of 200 and the given body
func OkResponseWithBody(body []byte) Response {
	return NewResponseWithBody(200, body)
}

// BadRequestResponse creates a new response with a status code of 400
func BadRequestResponse() Response {
	return NewResponse(400)
}

// BadRequestResponseWithBody creates a new response with a status code of 400 and the given body
func BadRequestResponseWithBody(body []byte) Response {
	return NewResponseWithBody(400, body)
}

// NotFoundResponse creates a new response with a status code of 404
func NotFoundResponse() Response {
	return NewResponse(404)
}

// NotFoundResponseWithBody creates a new response with a status code of 404 and the given body
func NotFoundResponseWithBody(body []byte) Response {
	return NewResponseWithBody(404, body)
}

// InternalErrorResponse creates a new response with a status code of 500
func InternalErrorResponse() Response {
	return NewResponse(500)
}

// InternalErrorResponseWithBody creates a new response with a status code of 500 and the given body
func InternalErrorResponseWithBody(body []byte) Response {
	return NewResponseWithBody(500, body)
}
