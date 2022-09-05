package handler

import (
	"strings"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"

	log "github.com/sirupsen/logrus"
)

const (
	Bullets         = 1
	GunFireMessage  = "💀🔫"
	GunClickMessage = "😌🔫"
)

type (
	Gun struct {
		C [6]bool
		N int
	}

	RouletteHandler struct {
		config bot.Config
		Name   string
	}
)

var (
	gun *Gun
)

func init() {
	gun = NewGun()
}

func NewGun() *Gun {
	return new(Gun)
}

func (g *Gun) Load(n int) {
	g.N = 0
	for i := 1; i <= n; {
		x := lib.RandInt(0, len(g.C)-1)
		if g.C[x] == false {
			g.C[x] = true
			i++
		} else {
			continue
		}
	}
}

func (g *Gun) Fire() bool {
	if g.C[g.N] {
		g.C[g.N] = false
		g.N++
		return true
	}

	g.N++
	return false
}

func (g *Gun) IsEmpty() bool {
	for _, v := range g.C {
		if v == true {
			return false
		}
	}

	return true
}

func NewRouletteHandler(s string) *RouletteHandler {
	return &RouletteHandler{Name: s}
}

func (h *RouletteHandler) SetConfig(config bot.Config) {
	h.config = config
}

func (h *RouletteHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!roulette") {
		return
	}

	if gun.IsEmpty() {
		gun.Load(Bullets)
		log.Debugf("reloading gun: %+v\n", gun)
	}

	log.Debugf("firing gun: %+v\n", gun)
	if gun.Fire() {
		s.ChannelMessageSend(m.ChannelID, GunFireMessage)
	} else {
		s.ChannelMessageSend(m.ChannelID, GunClickMessage)
	}
}
