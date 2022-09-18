package rps

import (
	"fmt"
	"strings"

	"git.kill0.net/chill9/beepboop/lib"
)

type (
	Game struct {
		rules    [][]string
		emojiMap map[string]string
	}
)

type InvalidChoiceError struct {
	s string
}

func (e InvalidChoiceError) Error() string {
	return fmt.Sprintf("%q is an invalid choice", e.s)
}

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

func NewGame(rules [][]string, emojiMap map[string]string) *Game {
	return &Game{rules: rules, emojiMap: emojiMap}
}

func (g *Game) Rand() string {
	return lib.MapRandKey(g.emojiMap)
}

func (g *Game) Play(c1, c2 string) (string, error) {
	var b strings.Builder

	if !g.Valid(c1) {
		return "", InvalidChoiceError{s: c1}
	}
	if !g.Valid(c2) {
		return "", InvalidChoiceError{s: c2}
	}

	fmt.Fprintf(&b, "%s v %s: ", g.emojiMap[c1], g.emojiMap[c2])

	for _, rule := range g.rules {
		verb := rule[2]

		if c1 == c2 {
			fmt.Fprintf(&b, "draw")
			return b.String(), nil
		}

		if c1 == rule[0] && c2 == rule[1] {
			fmt.Fprintf(&b, "%s %s %s", c1, verb, c2)
			return b.String(), nil
		} else if c2 == rule[0] && c1 == rule[1] {
			fmt.Fprintf(&b, "%s %s %s", c2, verb, c1)
			return b.String(), nil
		}
	}

	return b.String(), nil
}

func (g *Game) Valid(c string) bool {
	_, ok := g.emojiMap[c]
	return ok
}

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
