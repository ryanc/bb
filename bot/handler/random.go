package handler

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"

	log "github.com/sirupsen/logrus"
)

const (
	MaxDice         = 100
	MaxSides        = 100
	Bullets         = 1
	GunFireMessage  = "💀🔫"
	GunClickMessage = "😌🔫"
)

type (
	Roll struct {
		N, D, Sum int
		Rolls     []int
		S         string
	}

	Coin bool

	Gun struct {
		C [6]bool
		N int
	}

	CoinHandler struct {
		config bot.Config
	}
	RollHandler struct {
		config bot.Config
	}

	RouletteHandler struct {
		config bot.Config
	}
)

var (
	gun *Gun
)

func init() {
	gun = NewGun()
}

func NewRoll(n, d int) *Roll {
	r := new(Roll)
	r.N = n
	r.D = d
	r.S = fmt.Sprintf("%dd%d", r.N, r.D)
	return r
}

func ParseRoll(roll string) (*Roll, error) {
	var (
		dice []string
		err  error
		n, d int
	)

	match, _ := regexp.MatchString(`^(?:\d+)?d\d+$`, roll)

	if !match {
		return nil, errors.New("invalid roll, use `<n>d<sides>` e.g. `4d6`")
	}

	dice = strings.Split(roll, "d")

	if dice[0] == "" {
		n = 1
	} else {
		n, err = strconv.Atoi(dice[0])
		if err != nil {
			return nil, err
		}
	}

	d, err = strconv.Atoi(dice[1])
	if err != nil {
		return nil, err
	}

	if n > MaxDice || d > MaxSides {
		return nil, fmt.Errorf("invalid roll, n must be <= %d and sides must be <= %d", MaxDice, MaxSides)
	}

	return NewRoll(n, d), nil
}

func (r *Roll) RollDice() {
	for i := 1; i <= r.N; i++ {
		roll := lib.RandInt(1, r.D)
		r.Rolls = append(r.Rolls, roll)
		r.Sum += roll
	}
}

func (c *Coin) Flip() bool {
	*c = Coin(lib.Itob(lib.RandInt(0, 1)))
	return bool(*c)
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

func NewRollHandler() *RollHandler {
	return new(RollHandler)
}

func (h *RollHandler) SetConfig(config bot.Config) {
	h.config = config
}

func (h *RollHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		err       error
		msg, roll string
		r         *Roll
	)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!roll") {
		return
	}

	x := strings.Split(m.Content, " ")

	if len(x) != 2 {
		s.ChannelMessageSend(m.ChannelID, "help: `!roll <n>d<s>`")
		return
	}

	roll = x[1]

	r, err = ParseRoll(roll)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	r.RollDice()
	log.Debugf("rolled dice: %+v", r)

	msg = fmt.Sprintf("🎲 %s = %d", lib.JoinInt(r.Rolls, " + "), r.Sum)

	s.ChannelMessageSend(m.ChannelID, msg)
}

func NewRouletteHandler() *RouletteHandler {
	return new(RouletteHandler)
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

func NewCoinHandler() *CoinHandler {
	return new(CoinHandler)
}

func (h *CoinHandler) SetConfig(config bot.Config) {
	h.config = config
}

func (h *CoinHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		c   Coin
		msg string
	)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!coin") {
		return
	}

	if c.Flip() {
		msg = "heads"
	} else {
		msg = "tails"
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}
