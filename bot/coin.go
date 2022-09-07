package bot

import (
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
)

type Coin bool

func (c *Coin) Flip() bool {
	*c = Coin(lib.Itob(lib.RandInt(0, 1)))
	return bool(*c)
}

func (b *Bot) CoinCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		var (
			c   Coin
			msg string
		)

		if c.Flip() {
			msg = "heads"
		} else {
			msg = "tails"
		}

		b.Session.ChannelMessageSend(m.ChannelID, msg)
		return nil
	}
}
