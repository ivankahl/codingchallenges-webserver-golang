package webserver

import (
	"fmt"
	"regexp"
	"strings"
)

type Path struct {
	regex *regexp.Regexp
}

func (p Path) Matches(s string) bool {
	return p.regex.MatchString(s)
}

func RegexPath(regex *regexp.Regexp) Path {
	return Path{regex: regex}
}

func AnyPath() Path {
	return RegexPath(regexp.MustCompile("^.+$"))
}

func StringPath(path string) Path {
	regex := regexp.MustCompile(fmt.Sprintf("^%v\\/{0,1}$", strings.ReplaceAll(path, "/", "\\/")))

	return Path{regex: regex}
}
