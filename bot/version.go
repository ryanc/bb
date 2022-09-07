package bot

import (
	"fmt"
	"runtime"

	"github.com/bwmarrin/discordgo"
)

const (
	SourceURI = "https://git.kill0.net/chill9/bb"
)

func (b *Bot) VersionCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
			"go version: %s\nplatform: %s\nos: %s\nsource: %s\n",
			runtime.Version(),
			runtime.GOARCH,
			runtime.GOOS,
			SourceURI,
		))
		return nil
	}
}
