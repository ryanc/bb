package bot

import (
	"math/rand"
	"strings"

	"git.kill0.net/chill9/beepboop/lib"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type (
	ReactionHandler struct {
		Config Config
	}
)

func NewReactionHandler() *ReactionHandler {
	return new(ReactionHandler)
}

func (h *ReactionHandler) SetConfig(config Config) {
	h.Config = config
}

func (h *ReactionHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	emojis := h.Config.Handler.Reaction.Emojis
	channels := h.Config.Handler.Reaction.Channels

	if len(emojis) == 0 {
		log.Warning("emoji list is empty")
		return
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Fatalf("unable to get channel name: %v", err)
	}

	if len(channels) > 0 && !lib.Contains(channels, channel.Name) {
		return
	}

	for _, a := range m.Attachments {
		if strings.HasPrefix(a.ContentType, "image/") {
			for i := 1; i <= lib.RandInt(1, len(emojis)); i++ {
				r := emojis[rand.Intn(len(emojis))]
				s.MessageReactionAdd(m.ChannelID, m.ID, r)
			}
		}
	}

	for range m.Embeds {
		for i := 1; i <= lib.RandInt(1, len(emojis)); i++ {
			r := emojis[rand.Intn(len(emojis))]
			s.MessageReactionAdd(m.ChannelID, m.ID, r)
		}
	}
}
