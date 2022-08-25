package handler

import (
	"strings"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
)

type (
	Coin bool

	CoinHandler struct {
		config bot.Config
	}
)

func (c *Coin) Flip() bool {
	*c = Coin(lib.Itob(lib.RandInt(0, 1)))
	return bool(*c)
}

func NewCoinHandler() *CoinHandler {
	return new(CoinHandler)
}

func (h *CoinHandler) SetConfig(config bot.Config) {
	h.config = config
}

func (h *CoinHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		c   Coin
		msg string
	)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!coin") {
		return
	}

	if c.Flip() {
		msg = "heads"
	} else {
		msg = "tails"
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}
