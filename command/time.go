package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

const ()

var ()

type (
	TimeHandler struct{}
)

func NewTimeHandler() *TimeHandler {
	var h *TimeHandler = new(TimeHandler)
	return h
}

func (h *TimeHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var tz string

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!time") {
		return
	}

	x := strings.SplitN(m.Content, " ", 2)

	if len(x) != 2 {
		s.ChannelMessageSend(m.ChannelID, "help: `!time TIMEZONE`")
		return
	}

	tz = x[1]

	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Warnf("failed to load location: %s", err)
		return
	}

	now := time.Now()

	s.ChannelMessageSend(m.ChannelID, fmt.Sprint(now.In(loc)))
}
