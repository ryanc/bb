package commands

import (
	"fmt"
	"runtime"

	"git.kill0.net/chill9/beepboop/bot"
)

const (
	SourceURI = "https://git.kill0.net/chill9/bb"
)

func VersionCommand(cmd *bot.Command, args []string) error {
	cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, fmt.Sprintf(
		"go version: %s\nplatform: %s\nos: %s\nsource: %s\n",
		runtime.Version(),
		runtime.GOARCH,
		runtime.GOOS,
		SourceURI,
	))
	return nil
}
