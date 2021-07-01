package web

import (
	"fmt"
	"math/rand"
	"regexp"
)

const (
	randomInt  = "{random.int}"
	randomStr  = "{random.str}"
	randomLong = "{random.long}"
)

func replaceAllInt(s string) string {
	num := rand.Uint32()
	r := regexp.MustCompile("^(.*?){random.int}(.*)$")
	replaceStr := fmt.Sprintf("${1}%d", num)
	return r.ReplaceAllString(s, replaceStr)
}

func replaceRegexp(body string) string {
	body = replaceAllInt(body)
	return body
}
