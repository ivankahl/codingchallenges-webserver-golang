package webserver

import (
	"regexp"
	"testing"
)

func TestPath_MatchesValidPath(t *testing.T) {
	path := Path{
		regex: regexp.MustCompile("^/hello/world$"),
	}

	if !path.Matches("/hello/world") {
		t.Fatalf("Expected path to match /hello/world")
	}
}

func TestPath_MatchesInvalidPath(t *testing.T) {
	path := Path{
		regex: regexp.MustCompile("^/hello/world$"),
	}

	if path.Matches("/hello") {
		t.Fatalf("Expected path to not match /hello")
	}
}

func TestRegexPath(t *testing.T) {
	regex := regexp.MustCompile("^/hello/world$")
	path := RegexPath(regex)

	if !path.Matches("/hello/world") {
		t.Fatalf("Expected path to match /hello/world")
	}
}

func TestAnyPath(t *testing.T) {
	path := AnyPath()

	if !path.Matches("/hello/world") {
		t.Fatalf("Expected path to match /hello/world")
	}
}

func TestStringPath(t *testing.T) {
	path := StringPath("/hello/world")

	if !path.Matches("/hello/world") {
		t.Fatalf("Expected path to match /hello/world")
	}
}
