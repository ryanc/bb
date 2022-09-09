package bot

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	Rock = iota
	Paper
	Scissors
)

var (
	rpsVictoryMap map[int]int = map[int]int{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
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

func RpsRand() int {
	n := rand.Intn(len(rpsVictoryMap))
	i := 0
	for _, v := range rpsChoiceMap {
		if i == n {
			return v
		}
		i++
	}
	panic("unreachable")
}

func (b *Bot) RpsCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		var (
			botChoice, playerChoice int
			botEmoji, playerEmoji   string
			c                       string
		)

		if len(args) != 1 {
			b.Session.ChannelMessageSend(m.ChannelID, "help: `!rps (rock | paper | scissors)`")
			return nil
		}

		c = strings.ToLower(args[0])

		playerChoice, ok := rpsChoiceMap[c]
		if !ok {
			b.Session.ChannelMessageSend(m.ChannelID, "help: `!rps (rock | paper | scissors)`")
		}

		botChoice = RpsRand()
		botEmoji = rpsEmojiMap[botChoice]
		playerEmoji = rpsEmojiMap[playerChoice]

		if botChoice == playerChoice {
			b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s v %s: draw", botEmoji, playerEmoji))
			return nil
		} else if rpsVictoryMap[botChoice] == playerChoice {
			b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s v %s: bot wins", botEmoji, playerEmoji))
			return nil
		}

		b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s v %s: you win", botEmoji, playerEmoji))
		return nil
	}
}
