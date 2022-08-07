package command

import (
	"fmt"
	"runtime"

	"github.com/bwmarrin/discordgo"
)

const (
	SourceURI = "https://git.kill0.net/chill9/bb"
)

type (
	VersionHandler struct {
		config Config
		Name   string
	}
)

func NewVersionHandler(s string) *VersionHandler {
	h := new(VersionHandler)
	h.Name = s
	return h
}

func (h *VersionHandler) SetConfig(config Config) {
	h.config = config
}

func (h *VersionHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !HasCommand(m.Content, h.config.Prefix, h.Name) {
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
		"go version: %s\nplatform: %s\nos: %s\nsource: %s\n",
		runtime.Version(),
		runtime.GOARCH,
		runtime.GOOS,
		SourceURI,
	))
}
