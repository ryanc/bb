package command

import (
	"strings"

	"git.kill0.net/chill9/beepboop/lib/rps"
	"github.com/bwmarrin/discordgo"
)

func (h *Handlers) Rpsls(args []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	if len(args) != 1 {
		s.ChannelMessageSend(
			m.ChannelID, "help: `!rps (rock | paper | scissors | lizard | spock)`",
		)
		return nil
	}

	pc := strings.ToLower(args[0]) // player's choice

	g := rps.NewGame(rps.RulesRpsls, rps.EmojiMapRpsls)

	bc := g.Rand() // bot's choice

	out, err := g.Play(bc, pc)
	if _, ok := err.(rps.InvalidChoiceError); ok {
		s.ChannelMessageSend(
			m.ChannelID, "help: `!rps (rock | paper | scissors | lizard | spock)`",
		)
	}

	s.ChannelMessageSend(m.ChannelID, out)

	return nil
}
