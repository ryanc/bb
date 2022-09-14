package bot

import (
	"fmt"
	"strings"

	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
)

var (
	rpslsVictoryMap map[int][]int = map[int][]int{
		Rock:     {Scissors, Lizard},
		Paper:    {Rock, Spock},
		Scissors: {Paper, Lizard},
		Lizard:   {Paper, Spock},
		Spock:    {Scissors, Rock},
	}

	rpslsChoiceMap map[string]int = map[string]int{
		"rock":     Rock,
		"paper":    Paper,
		"scissors": Scissors,
		"lizard":   Lizard,
		"spock":    Spock,
	}

	rpslsEmojiMap map[int]string = map[int]string{
		Rock:     "🪨️",
		Paper:    "📝",
		Scissors: "✂️",
		Lizard:   "🦎",
		Spock:    "🖖",
	}
)

func (b *Bot) RpslsCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		var (
			bc, pc int
			be, pe string
			c      string
		)

		if len(args) != 1 {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rpsls (rock | paper | scissors | lizard | spock )`",
			)
			return nil
		}

		c = strings.ToLower(args[0])

		pc, ok := rpslsChoiceMap[c] // player's choice
		if !ok {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rpsls (rock | paper | scissors | lizard | spock)`",
			)
		}

		bc = lib.MapRand(rpslsChoiceMap) // bot's choice
		pe = rpslsEmojiMap[pc]           // player's emoji
		be = rpslsEmojiMap[bc]           // bot's emoji

		if bc == pc {
			b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
				"%s v %s: draw", be, pe,
			))
			return nil
		} else if lib.Contains(rpslsVictoryMap[bc], pc) {
			b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
				"%s v %s: %s wins", be, pe, lib.MapKey(rpslsChoiceMap, bc),
			))
			return nil
		}

		b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
			"%s v %s: %s wins", be, pe, lib.MapKey(rpslsChoiceMap, pc),
		))
		return nil
	}
}
