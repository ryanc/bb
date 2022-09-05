package handler

import (
	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type (
	PingHandler struct {
		config bot.Config
		Name   string
	}
)

func NewPingHandler(s string) *PingHandler {
	return &PingHandler{Name: s}
}

func (h *PingHandler) SetConfig(config bot.Config) {
	h.config = config
}

func (h *PingHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !lib.ContainsCommand(m.Content, h.config.Prefix, h.Name) {
		return
	}

	log.Debug("received ping")

	s.ChannelMessageSend(m.ChannelID, "pong")
}
