package bot

import (
	"fmt"
	"strings"

	"git.kill0.net/chill9/beepboop/lib"
	"git.kill0.net/chill9/beepboop/lib/rps"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RpslsCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		var (
			bc, pc, be, pe string
		)

		if len(args) != 1 {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rpsls (rock | paper | scissors | lizard | spock)`",
			)
			return nil
		}

		pc = strings.ToLower(args[0])

		_, ok := rps.EmojiMapRpsls[pc] // player's choice
		if !ok {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rpsls (rock | paper | scissors | lizard | spock)`",
			)
			return nil
		}

		bc = lib.MapRandKey(rps.EmojiMapRpsls) // bot's choice
		pe = rps.EmojiMapRpsls[pc]             // player's emoji
		be = rps.EmojiMapRpsls[bc]             // bot's emoji

		b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
			"%s v %s: %s", pe, be, rps.Play(rps.RulesRpsls, bc, pc),
		))
		return nil
	}
}
