package bot

import (
	"fmt"
	"strings"

	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
)

const (
	Rock = iota
	Paper
	Scissors
	Lizard
	Spock
)

var (
	rpsVictoryMap map[int][]int = map[int][]int{
		Rock:     {Scissors},
		Paper:    {Rock},
		Scissors: {Paper},
	}

	rpsChoiceMap map[string]int = map[string]int{
		"rock":     Rock,
		"paper":    Paper,
		"scissors": Scissors,
	}

	rpsEmojiMap map[int]string = map[int]string{
		Rock:     "🪨️",
		Paper:    "📝",
		Scissors: "✂️",
	}
)

func (b *Bot) RpsCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		var (
			bc, pc int
			be, pe string
			c      string
		)

		if len(args) != 1 {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rps (rock | paper | scissors)`",
			)
			return nil
		}

		c = strings.ToLower(args[0])

		pc, ok := rpsChoiceMap[c] // player's choice
		if !ok {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rps (rock | paper | scissors)`",
			)
		}

		bc = lib.MapRand(rpsChoiceMap) // bot's choice
		pe = rpsEmojiMap[pc]           // player's emoji
		be = rpsEmojiMap[bc]           // bot's emoji

		if bc == pc {
			b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
				"%s v %s: draw", be, pe,
			))
			return nil
		} else if lib.Contains(rpsVictoryMap[bc], pc) {
			b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
				"%s v %s: %s wins", be, pe, lib.MapKey(rpsChoiceMap, bc),
			))
			return nil
		}

		b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
			"%s v %s: %s wins", be, pe, lib.MapKey(rpsChoiceMap, pc),
		))
		return nil
	}
}
