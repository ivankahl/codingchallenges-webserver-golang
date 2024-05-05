package webserver

import (
	"errors"
	"strings"
)

var ErrInvalidMethod = errors.New("invalid method")

type Method string

const (
	MethodGet     Method = "GET"
	MethodHead           = "HEAD"
	MethodPost           = "POST"
	MethodPut            = "PUT"
	MethodDelete         = "DELETE"
	MethodConnect        = "CONNECT"
	MethodOptions        = "OPTIONS"
	MethodTrace          = "TRACE"
	MethodPatch          = "PATCH"
	MethodAny            = "*"
)

func methodFromString(s string) (Method, error) {
	switch strings.ToUpper(s) {
	case string(MethodGet):
		return MethodGet, nil
	case MethodHead:
		return MethodHead, nil
	case MethodPost:
		return MethodPost, nil
	case MethodPut:
		return MethodPut, nil
	case MethodDelete:
		return MethodDelete, nil
	case MethodConnect:
		return MethodConnect, nil
	case MethodOptions:
		return MethodOptions, nil
	case MethodTrace:
		return MethodTrace, nil
	case MethodPatch:
		return MethodPatch, nil
	case MethodAny:
		return MethodAny, nil
	default:
		return "", ErrInvalidMethod
	}
}
