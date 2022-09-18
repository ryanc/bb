package bot

import (
	"strings"

	"git.kill0.net/chill9/beepboop/lib/rps"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RpslsCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		if len(args) != 1 {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rps (rock | paper | scissors | lizard | spock)`",
			)
			return nil
		}

		pc := strings.ToLower(args[0]) // player's choice

		g := rps.NewGame(rps.RulesRpsls, rps.EmojiMapRpsls)

		bc := g.Rand() // bot's choice

		s, err := g.Play(bc, pc)
		if _, ok := err.(rps.InvalidChoiceError); ok {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rps (rock | paper | scissors | lizard | spock)`",
			)
		}

		b.Session.ChannelMessageSend(m.ChannelID, s)

		return nil
	}
}
