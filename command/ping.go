package command

import (
	"strings"

	"git.kill0.net/chill9/beepboop/bot"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type (
	PingHandler struct {
		config bot.Config
	}
)

func NewPingHandler() *PingHandler {
	return new(PingHandler)
}

func (h *PingHandler) SetConfig(config bot.Config) {
	h.config = config
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
