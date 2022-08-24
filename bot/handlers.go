package bot

import (
	"github.com/bwmarrin/discordgo"
)

type MessageCreateHandler interface {
	Handle(*discordgo.Session, *discordgo.MessageCreate)
	SetConfig(Config)
}
