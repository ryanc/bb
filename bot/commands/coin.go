package commands

import (
	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"
)

type Coin bool

func (c *Coin) Flip() bool {
	*c = Coin(lib.Itob(lib.RandInt(0, 1)))
	return bool(*c)
}

func CoinCommand(cmd *bot.Command, args []string) error {
	var (
		c   Coin
		msg string
	)

	if c.Flip() {
		msg = "heads"
	} else {
		msg = "tails"
	}

	cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, msg)
	return nil
}
