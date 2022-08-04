package command

import "github.com/bwmarrin/discordgo"

type CommandHandler interface {
	Handle(s *discordgo.Session, m *discordgo.MessageCreate)
	SetConfig(config Config)
}
