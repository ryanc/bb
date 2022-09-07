package bot

import "github.com/bwmarrin/discordgo"

type Bot struct {
	Session *discordgo.Session
	Config  Config
}

func NewBot(s *discordgo.Session, config Config) *Bot {
	return &Bot{Session: s, Config: config}
}

func (b *Bot) RegisterCommands() {
	AddCommand(&Command{
		Name: "coin",
		Func: b.CoinCommand(),
	})
	AddCommand(&Command{
		Name:  "deal",
		Func:  b.DealCommand(),
		NArgs: 1,
	})
	AddCommand(&Command{
		Name: "ping",
		Func: b.PingCommand(),
	})
	AddCommand(&Command{
		Name:  "roll",
		Func:  b.RollCommand(),
		NArgs: 1,
	})
	AddCommand(&Command{
		Name:  "time",
		Func:  b.TimeCommand(),
		NArgs: 1,
	})
	AddCommand(&Command{
		Name: "version",
		Func: b.VersionCommand(),
	})
	AddCommand(&Command{
		Name:  "weather",
		Func:  b.WeatherCommand(),
		NArgs: 1,
	})
}
