package command

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

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

	CoinHandler struct{}
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
		roll := RandInt(1, r.D)
		r.Rolls = append(r.Rolls, roll)
		r.Sum += roll
	}
}

func (c *Coin) Flip() bool {
	*c = Coin(Itob(RandInt(0, 1)))
	return bool(*c)
}

func NewGun() *Gun {
	return new(Gun)
}

func (g *Gun) Load(n int) {
	g.N = 0
	for i := 1; i <= n; {
		x := RandInt(0, len(g.C)-1)
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

func RollHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	msg = fmt.Sprintf("🎲 %s = %d", JoinInt(r.Rolls, " + "), r.Sum)

	s.ChannelMessageSend(m.ChannelID, msg)
}

func RouletteHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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

func RandInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func JoinInt(a []int, sep string) string {
	var b []string

	b = make([]string, len(a))

	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, sep)
}

func SumInt(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

func Itob(v int) bool {
	if v == 1 {
		return true
	}

	return false
}
