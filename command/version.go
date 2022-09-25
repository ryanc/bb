package command

import (
	"fmt"
	"runtime"

	"github.com/bwmarrin/discordgo"
)

const (
	SourceURI = "https://git.kill0.net/chill9/bb"
)

func (h *Handlers) Version(args []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
		"go version: %s\nplatform: %s\nos: %s\nsource: %s\n",
		runtime.Version(),
		runtime.GOARCH,
		runtime.GOOS,
		SourceURI,
	))
	return nil
}
