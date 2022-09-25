package command

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) Time(args []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	var (
		t  time.Time
		tz string
	)

	now := time.Now()

	if len(args) == 1 {
		tz = args[0]
		loc, err := time.LoadLocation(tz)
		if err != nil {
			log.Warnf("failed to load location: %s", err)
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return nil
		}
		t = now.In(loc)
	} else {
		t = now
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprint(t))
	return nil
}
