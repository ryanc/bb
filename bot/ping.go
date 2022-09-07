package bot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) PingCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		b.Session.ChannelMessageSend(m.ChannelID, "pong")
		return nil
	}
}
