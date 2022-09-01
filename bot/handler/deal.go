package handler

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type (
	DealHandler struct {
		config bot.Config
		Name   string
	}

	Card string

	Deck [52]Card
)

var deck Deck = Deck{
	"2♣", "3♣", "4♣", "5♣", "6♣", "7♣", "8♣", "9♣", "10♣", "J♣", "Q♣", "K♣", "A♣",
	"2♦", "3♦", "4♦", "5♦", "6♦", "7♦", "8♦", "9♦", "10♦", "J♦", "Q♦", "K♦", "A♦",
	"2♥", "3♥", "4♥", "5♥", "6♥", "7♥", "8♥", "9♥", "10♥", "J♥", "Q♥", "K♥", "A♥",
	"2♠", "3♠", "4♠", "5♠", "6♠", "7♠", "8♠", "9♠", "10♠", "J♠", "Q♠", "K♠", "A♠",
}

func (d *Deck) Deal(n int) ([]Card, error) {
	var (
		hand []Card
		err  error
	)

	if n < 1 {
		err = errors.New("number cannot be less than 1")
		return hand, err
	}

	if n > len(d) {
		err = errors.New("number is greater than cards in the deck")
		return hand, err
	}

	hand = deck[0:n]

	return hand, err
}

func NewDealHandler(s string) *DealHandler {
	h := new(DealHandler)
	h.Name = s
	return h
}

func (h *DealHandler) SetConfig(config bot.Config) {
	h.config = config
}

func (h *DealHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !lib.HasCommand(m.Content, h.config.Prefix, h.Name) {
		return
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	log.Debugf("%+v", deck)

	_, args := lib.SplitCommandAndArgs(m.Content, h.config.Prefix)

	if len(args) != 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("help: `!%s <n>`", h.Name))
		return
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		log.Errorf("failed to convert string to int: %s", err)
	}

	hand, err := deck.Deal(n)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error: %s\n", err))
		return
	}

	s.ChannelMessageSend(m.ChannelID, JoinCards(hand, " "))
}

func JoinCards(h []Card, sep string) string {
	var b []string

	b = make([]string, len(h))

	for i, v := range h {
		b[i] = string(v)
	}

	return strings.Join(b, sep)
}
