package bot

import (
	"github.com/bwmarrin/discordgo"
)

type MessageCreateHandler interface {
	Handle(s *discordgo.Session, m *discordgo.MessageCreate)
	SetConfig(config Config)
}
