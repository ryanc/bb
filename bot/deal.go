package bot

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type (
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

func (b *Bot) DealCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		rand.Shuffle(len(deck), func(i, j int) {
			deck[i], deck[j] = deck[j], deck[i]
		})

		log.Debugf("%+v", deck)

		if len(args) != 1 {
			b.Session.ChannelMessageSend(m.ChannelID, "help: `!deal <n>`")
			return nil
		}

		n, err := strconv.Atoi(args[0])
		if err != nil {
			log.Errorf("failed to convert string to int: %s", err)
		}

		hand, err := deck.Deal(n)
		if err != nil {
			b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error: %s\n", err))
			return nil
		}

		b.Session.ChannelMessageSend(m.ChannelID, JoinCards(hand, " "))
		return nil
	}
}

func JoinCards(h []Card, sep string) string {
	var b []string

	b = make([]string, len(h))

	for i, v := range h {
		b[i] = string(v)
	}

	return strings.Join(b, sep)
}
