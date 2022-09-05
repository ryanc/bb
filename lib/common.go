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

func HasCommand(s, prefix string) bool {
	s = strings.TrimSpace(s)

	if len(s) == 0 || len(prefix) == 0 {
		return false
	}

	if !strings.HasPrefix(s, prefix) {
		return false
	}

	// remove the command prefix
	s = s[len(prefix):]

	// multiple assignment trick
	cmd, _ := func() (string, string) {
		x := strings.SplitN(s, " ", 2)
		if len(x) > 1 {
			return x[0], x[1]
		}

		return x[0], ""
	}()

	return len(cmd) > 0
}

func ContainsCommand(s, prefix, cmd string) bool {
	s = strings.TrimSpace(s)

	args := strings.Split(s, " ")
	s = args[0]

	// a command cannot be less than two characters e.g. !x
	if len(s) < 2 {
		return false
	}

	if string(s[0]) != prefix {
		return false
	}

	if strings.HasPrefix(s[1:], cmd) {
		return true
	}

	return false
}

func SplitCommandAndArgs(s, prefix string) (cmd string, args []string) {
	s = strings.TrimSpace(s)

	x := strings.Split(s, " ")

	if len(x) > 1 {
		args = x[1:]
	}

	cmd = x[0]

	if strings.Index(s, prefix) == 0 {
		cmd = cmd[len(prefix):]
	}

	return cmd, args
}
