package lib

import (
	"reflect"
	"testing"
)

func TestContainsCommand(t *testing.T) {
	tables := []struct {
		s      string
		prefix string
		cmd    string
		r      bool
	}{
		{"!command x y", "!", "command", true},
		{"#command x y", "#", "command", true},
		{"command x y", "!", "comamnd", false},
		{"", "", "", false},
		{"!", "!", "", false},
		{"! x y", "!", "", false},
	}

	for _, table := range tables {
		r := ContainsCommand(table.s, table.prefix, table.cmd)
		if r != table.r {
			t.Errorf("ContainsCommand(%q, %q, %q), got: %t, want: %t",
				table.s, table.prefix, table.cmd, r, table.r,
			)
		}
	}
}

func TestHasCommandCommand(t *testing.T) {
	tables := []struct {
		s      string
		prefix string
		want   bool
	}{
		{"!command", "!", true},
		{"!command x y", "!", true},
		{"!c x y", "!", true},
		{"! x y", "!", false},
		{"hey guy", "!", false},
		{"hey", "!", false},
		{"hey", "", false},
		{"", "!", false},
		{"", "", false},
	}

	for _, table := range tables {
		if got, want := HasCommand(table.s, table.prefix), table.want; got != want {
			t.Errorf(
				"s: %s, prefix: %s, got: %t, want: %t",
				table.s, table.prefix, got, want,
			)
		}
	}
}

func TestSplitCommandAndArg(t *testing.T) {
	tables := []struct {
		s       string
		prefix  string
		wantCmd string
		wantArg string
	}{
		{"!command x y", "!", "command", "x y"},
		{"!command", "!", "command", ""},
		{"hey man", "!", "", ""},
	}

	for _, table := range tables {
		gotCmd, gotArg := SplitCommandAndArg(table.s, table.prefix)
		if gotCmd != table.wantCmd {
			t.Errorf("got: %s, want: %s", gotCmd, table.wantCmd)
		}
		if gotArg != table.wantArg {
			t.Errorf("got: %+v, want: %+v", gotArg, table.wantArg)
		}
	}
}

func TestSplitCommandAndArgs(t *testing.T) {
	tables := []struct {
		s        string
		prefix   string
		n        int
		wantCmd  string
		wantArgs []string
	}{
		{"!command x y", "!", 2, "command", []string{"x", "y"}},
		{"!command x y z", "!", 2, "command", []string{"x", "y z"}},
		{"!command", "!", 1, "command", []string{""}},
		{"hey man", "!", 1, "", []string{""}},
	}
	for _, table := range tables {
		gotCmd, gotArgs := SplitCommandAndArgs(table.s, table.prefix, table.n)
		if gotCmd != table.wantCmd {
			t.Errorf("got: %s, want: %s", gotCmd, table.wantCmd)
		}
		if !reflect.DeepEqual(gotArgs, table.wantArgs) {
			t.Errorf("got: %+v, want: %+v", gotArgs, table.wantArgs)
		}
	}
}

func TestSplitArgs(t *testing.T) {
	tables := []struct {
		s    string
		n    int
		want []string
	}{
		{"a b c", 0, []string{"a", "b", "c"}},
		{"a b c", 1, []string{"a b c"}},
		{"a b c", 2, []string{"a", "b c"}},
		{"a b c", 3, []string{"a", "b", "c"}},
		{"a b c", 4, []string{"a", "b", "c"}},
	}
	for _, table := range tables {
		if got, want := SplitArgs(table.s, table.n), table.want; !reflect.DeepEqual(got, want) {
			t.Errorf("got: %#v, want: %#v", got, want)
		}
	}
}
