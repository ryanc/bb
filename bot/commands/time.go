package commands

import (
	"fmt"
	"time"

	"git.kill0.net/chill9/beepboop/bot"
	log "github.com/sirupsen/logrus"
)

func TimeCommand(cmd *bot.Command, args []string) error {
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
			cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, err.Error())
			return nil
		}
		t = now.In(loc)
	} else {
		t = now
	}

	cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, fmt.Sprint(t))
	return nil
}
