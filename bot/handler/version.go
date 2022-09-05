package handler

import (
	"fmt"
	"runtime"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
)

const (
	SourceURI = "https://git.kill0.net/chill9/bb"
)

type (
	VersionHandler struct {
		config bot.Config
		Name   string
	}
)

func NewVersionHandler(s string) *VersionHandler {
	return &VersionHandler{Name: s}
}

func (h *VersionHandler) SetConfig(config bot.Config) {
	h.config = config
}

func (h *VersionHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !lib.HasCommand(m.Content, h.config.Prefix, h.Name) {
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
