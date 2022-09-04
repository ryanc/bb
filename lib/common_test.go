package lib

import "testing"

func TestHasCommand(t *testing.T) {
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
		r := HasCommand(table.s, table.prefix, table.cmd)
		if r != table.r {
			t.Errorf("HasCommand(%q, %q, %q), got: %t, want: %t",
				table.s, table.prefix, table.cmd, r, table.r,
			)
		}
	}
}
