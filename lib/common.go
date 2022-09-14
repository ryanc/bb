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
	b := make([]string, len(a))

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
	return v == 1
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

func SplitCommandAndArg(s, prefix string) (cmd string, arg string) {
	s = strings.TrimSpace(s)

	if !strings.HasPrefix(s, prefix) {
		return
	}

	// remove the command prefix
	s = s[len(prefix):]

	// multiple assignment trick
	cmd, arg = func() (string, string) {
		x := strings.SplitN(s, " ", 2)
		if len(x) > 1 {
			return x[0], x[1]
		}
		return x[0], ""
	}()

	return cmd, arg
}

func SplitCommandAndArgs(s, prefix string, n int) (cmd string, args []string) {
	cmd, arg := SplitCommandAndArg(s, prefix)

	if arg == "" {
		return cmd, []string{}
	}

	if n == 0 {
		return cmd, strings.Split(arg, " ")
	}

	return cmd, strings.SplitN(arg, " ", n)
}

func SplitArgs(s string, n int) (args []string) {
	if s == "" {
		return []string{}
	}

	if n > 0 {
		args = strings.SplitN(s, " ", n)
	} else {
		args = strings.Split(s, " ")
	}
	return
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapKey[K comparable, V comparable](m map[K]V, v V) K {
	var r K
	for k := range m {
		if m[k] == v {
			return k
		}
	}
	return r
}
