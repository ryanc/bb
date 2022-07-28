package command

import (
	"math/rand"
	"net/url"
	"path"
	"strconv"
	"strings"
)

func RandInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
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
