package bot

import (
	"fmt"
	"strings"

	"git.kill0.net/chill9/beepboop/lib"
	"git.kill0.net/chill9/beepboop/lib/rps"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RpsCommand() CommandFunc {
	return func(args []string, m *discordgo.MessageCreate) error {
		var (
			bc, pc, be, pe string
		)

		if len(args) != 1 {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rps (rock | paper | scissors)`",
			)
			return nil
		}

		pc = strings.ToLower(args[0])

		_, ok := rps.EmojiMapRps[pc] // player's choice
		if !ok {
			b.Session.ChannelMessageSend(
				m.ChannelID, "help: `!rps (rock | paper | scissors)`",
			)
			return nil
		}

		bc = lib.MapRandKey(rps.EmojiMapRps) // bot's choice
		pe = rps.EmojiMapRps[pc]             // player's emoji
		be = rps.EmojiMapRps[bc]             // bot's emoji

		b.Session.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
			"%s v %s: %s", pe, be, rps.Play(rps.RulesRps, bc, pc),
		))
		return nil
	}
}
