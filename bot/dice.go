package bot

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

func (b *Bot) RollCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		var (
			err       error
			msg, roll string
			r         *Roll
		)

		roll = args[0]

		r, err = ParseRoll(roll)
		if err != nil {
			b.Session.ChannelMessageSend(m.ChannelID, err.Error())
			return nil
		}

		r.RollDice()
		log.Debugf("rolled dice: %+v", r)

		msg = fmt.Sprintf("🎲 %s = %d", lib.JoinInt(r.Rolls, " + "), r.Sum)

		b.Session.ChannelMessageSend(m.ChannelID, msg)
		return nil
	}
}
