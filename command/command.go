package command

import (
	"git.kill0.net/chill9/beepboop/bot"
	"github.com/bwmarrin/discordgo"
)

type CommandHandler interface {
	Handle(s *discordgo.Session, m *discordgo.MessageCreate)
	SetConfig(config bot.Config)
}
