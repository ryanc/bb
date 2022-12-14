package bot

import (
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

func init() {
	pflag.Bool("debug", false, "enable debug mode")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
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

func (b *Bot) RegisterHandlers() {
	b.Session.AddHandler(b.CommandHandler())
	b.Session.AddHandler(b.ReactionHandler())
}

func Run() error {
	initConfig()
	go reloadConfig()

	if err := lib.SeedMathRand(); err != nil {
		log.Warn(err)
	}

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

	if err = dg.Open(); err != nil {
		log.Fatalf("error opening connection: %v\n", err)
	}

	log.Info("The bot is now running. Press CTRL-C to exit.")

	defer dg.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	log.Info("Shutting down")

	return nil
}

func initConfig() {
	C = NewConfig()

	viper.SetEnvPrefix("BEEPBOOP")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("debug", false)

	viper.BindEnv("DISCORD_TOKEN")
	viper.BindEnv("OPEN_WEATHER_MAP_TOKEN")

	loadConfig()
}

func loadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("fatal error config file: %v", err)
		}
	}

	log.WithField("filename", viper.ConfigFileUsed()).Info(
		"loaded configuration file",
	)

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func reloadConfig() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP)
	for {
		<-sc

		loadConfig()
	}
}
