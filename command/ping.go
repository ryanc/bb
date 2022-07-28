package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	log "github.com/sirupsen/logrus"
)

type (
	PingHandler struct{}
)

func NewPingHandler() *PingHandler {
	return new(PingHandler)
}

func (h *PingHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!ping") {
		return
	}

	log.Debug("received ping")

	s.ChannelMessageSend(m.ChannelID, "pong")
}
