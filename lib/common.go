package lib

import (
	"net/url"
	"path"
	"strconv"
	"strings"
)

func Contains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func JoinInt(a []int, sep string) string {
	var b []string

	b = make([]string, len(a))

	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, sep)
}

func SumInt(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

func Itob(v int) bool {
	if v == 1 {
		return true
	}

	return false
}

func BuildURI(rawuri, rawpath string) string {
	u, _ := url.Parse(rawuri)
	u.Path = path.Join(u.Path, rawpath)
	return u.String()
}

func HasCommand(s, prefix, cmd string) bool {
	if len(s) < 2 {
		return false
	}

	if string(s[0]) != prefix {
		return false
	}

	if s[1:] == cmd {
		return true
	}

	return false
}
