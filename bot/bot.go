package bot

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var C Config

type (
	Bot struct {
		Session *discordgo.Session
		Config  Config
	}

	MessageHandler func(s *discordgo.Session, m *discordgo.MessageCreate)
)

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

func (b *Bot) RegisterHandlers() {
	b.Session.AddHandler(b.CommandHandler())
	b.Session.AddHandler(b.ReactionHandler())
}

func Run() error {
	setupConfig()

	lib.SeedMathRand()

	if C.DiscordToken == "" {
		log.Fatalf("Discord token is not set")
	}

	dg, err := discordgo.New(fmt.Sprintf("Bot %s", C.DiscordToken))
	if err != nil {
		log.Fatalf("error creating Discord session: %v\n", err)
	}

	b := NewBot(dg, C)
	b.RegisterHandlers()
	b.RegisterCommands()

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	err = dg.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v\n", err)
	}

	log.Info("The bot is now running. Press CTRL-C to exit.")

	defer dg.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Info("Shutting down")

	return nil
}

func setupConfig() {
	var err error

	C = NewConfig()

	flag.Bool("debug", false, "enable debug logging")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetEnvPrefix("BEEPBOOP")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()

	viper.BindEnv("DEBUG")
	viper.BindEnv("DISCORD_TOKEN")
	viper.BindEnv("OPEN_WEATHER_MAP_TOKEN")

	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		log.Fatalf("fatal error config file: %v", err)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
}
