package command

import (
	"github.com/bwmarrin/discordgo"
)

func (h *Handlers) Ping(args []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	s.ChannelMessageSend(m.ChannelID, "pong")
	return nil
}
