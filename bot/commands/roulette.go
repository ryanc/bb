package commands

import (
	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"

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

func RouletteCommand(cmd *bot.Command, args []string) error {
	if gun.IsEmpty() {
		gun.Load(Bullets)
		log.Debugf("reloading gun: %+v\n", gun)
	}

	log.Debugf("firing gun: %+v\n", gun)
	if gun.Fire() {
		cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, GunFireMessage)
	} else {
		cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, GunClickMessage)
	}
	return nil
}
