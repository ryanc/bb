package commands

import (
	"git.kill0.net/chill9/beepboop/bot"
)

func PingCommand(cmd *bot.Command, args []string) error {
	cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, "pong")
	return nil
}
