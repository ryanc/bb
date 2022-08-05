package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	SourceURI = "https://git.kill0.net/chill9/bb"
)

type (
	SourceHandler struct {
		config Config
	}
)

func NewSourceHandler() *SourceHandler {
	return new(SourceHandler)
}

func (h *SourceHandler) SetConfig(config Config) {
	h.config = config
}

func (h *SourceHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!source") {
		return
	}

	s.ChannelMessageSend(m.ChannelID, SourceURI)
}
