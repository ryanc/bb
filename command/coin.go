package command

import (
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
)

type Coin bool

func (c *Coin) Flip() bool {
	*c = Coin(lib.Itob(lib.RandInt(0, 1)))
	return bool(*c)
}

func (h *Handlers) Coin(args []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	var (
		c   Coin
		msg string
	)

	if c.Flip() {
		msg = "heads"
	} else {
		msg = "tails"
	}

	s.ChannelMessageSend(m.ChannelID, msg)
	return nil
}
