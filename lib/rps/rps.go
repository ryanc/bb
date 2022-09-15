package rps

import (
	"fmt"
)

type (
	Game struct {
		rules    [][]string
		emojiMap map[string]string
	}
)

var (
	RulesRps [][]string = [][]string{
		{"rock", "scissors", "crushes"},
		{"paper", "rock", "covers"},
		{"scissors", "paper", "cuts"},
	}

	EmojiMapRps map[string]string = map[string]string{
		"rock":     "🪨️",
		"paper":    "📝",
		"scissors": "✂️",
	}

	RulesRpsls [][]string = [][]string{
		{"rock", "scissors", "crushes"},
		{"rock", "lizard", "crushes"},
		{"paper", "rock", "covers"},
		{"paper", "spock", "disproves"},
		{"scissors", "paper", "cuts"},
		{"scissors", "lizard", "decapitates"},
		{"lizard", "paper", "eats"},
		{"lizard", "spock", "poisons"},
		{"spock", "scissors", "smashes"},
		{"spock", "rock", "vaporizes"},
	}

	EmojiMapRpsls map[string]string = map[string]string{
		"rock":     "🪨️",
		"paper":    "📝",
		"scissors": "✂️",
		"lizard":   "🦎",
		"spock":    "🖖",
	}
)

func Play(rules [][]string, c1, c2 string) string {
	for _, rule := range rules {
		if c1 == c2 {
			return "draw"
		}

		if c1 == rule[0] && c2 == rule[1] {
			return fmt.Sprintf("%s %s %s", c1, rule[2], c2)
		} else if c2 == rule[0] && c1 == rule[1] {
			return fmt.Sprintf("%s %s %s", c2, rule[2], c1)
		}
	}
	return ""
}
