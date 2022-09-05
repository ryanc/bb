package lib

import "testing"

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
