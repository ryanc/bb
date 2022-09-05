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
	MaxDice  = 100
	MaxSides = 100
)

type (
	Roll struct {
		N, D, Sum int
		Rolls     []int
		S         string
	}

	RollHandler struct {
		config bot.Config
		Name   string
	}
)

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

func NewRollHandler(s string) *RollHandler {
	return &RollHandler{Name: s}
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

	if !lib.HasCommand(m.Content, h.config.Prefix, h.Name) {
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
