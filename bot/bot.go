package bot

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"git.kill0.net/chill9/beepboop/command"
	"git.kill0.net/chill9/beepboop/config"
	"git.kill0.net/chill9/beepboop/handler"
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var C *config.Config

type (
	Bot struct {
		config   *config.Config
		session  *discordgo.Session
		commands map[string]*Command
	}

	Command struct {
		Name  string
		Func  CommandFunc
		NArgs int
	}

	CommandFunc func(args []string, s *discordgo.Session, m *discordgo.MessageCreate) error

	MessageHandler func(s *discordgo.Session, m *discordgo.MessageCreate)
)

func init() {
	pflag.Bool("debug", false, "enable debug mode")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}

func NewBot(config *config.Config, s *discordgo.Session) *Bot {
	return &Bot{
		session:  s,
		commands: make(map[string]*Command),
	}
}

func (b *Bot) AddHandler(handler interface{}) func() {
	return b.session.AddHandler(handler)
}

func (b *Bot) AddCommand(cmd *Command) {
	b.commands[cmd.Name] = cmd
}

func (b *Bot) GetCommand(name string) (*Command, bool) {
	cmd, ok := b.commands[name]
	return cmd, ok
}

func (b *Bot) CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !lib.HasCommand(m.Content, b.config.Prefix) {
		return
	}

	cmdName, arg := lib.SplitCommandAndArg(m.Content, b.config.Prefix)

	cmd, ok := b.GetCommand(cmdName)
	if !ok {
		return
	}

	args := lib.SplitArgs(arg, cmd.NArgs)

	if ok {
		log.Debugf("command: %v, args: %v, nargs: %d", cmd.Name, args, len(args))
		if err := cmd.Func(args, s, m); err != nil {
			log.Errorf("failed to execute command: %s", err)
		}

		return
	}

	log.Warnf("unknown command: %v, args: %v, nargs: %d", cmdName, args, len(args))
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("unknown command: %s", cmdName))
}

func (b *Bot) Init(h *handler.Handlers, ch *command.Handlers) {
	// Register handlers
	b.AddHandler(h.Reaction)

	// Register commands
	b.AddCommand(&Command{
		Name: "coin",
		Func: ch.Coin,
	})
	b.AddCommand(&Command{
		Name:  "deal",
		Func:  ch.Deal,
		NArgs: 1,
	})
	b.AddCommand(&Command{
		Name: "ping",
		Func: ch.Ping,
	})
	b.AddCommand(&Command{
		Name:  "roll",
		Func:  ch.Roll,
		NArgs: 1,
	})
	b.AddCommand(&Command{
		Name: "roulette",
		Func: ch.Roulette,
	})
	b.AddCommand(&Command{
		Name:  "rps",
		Func:  ch.Rps,
		NArgs: 1,
	})
	b.AddCommand(&Command{
		Name:  "rpsls",
		Func:  ch.Rpsls,
		NArgs: 1,
	})
	b.AddCommand(&Command{
		Name:  "time",
		Func:  ch.Time,
		NArgs: 1,
	})
	b.AddCommand(&Command{
		Name: "version",
		Func: ch.Version,
	})
	b.AddCommand(&Command{
		Name:  "weather",
		Func:  ch.Weather,
		NArgs: 1,
	})
}

func Run() error {
	initConfig()
	go reloadConfig()

	if err := lib.SeedMathRand(); err != nil {
		log.Warn(err)
	}

	if C.DiscordToken == "" {
		return errors.New("discord token not set")
	}

	dg, err := discordgo.New(fmt.Sprintf("Bot %s", C.DiscordToken))
	if err != nil {
		return fmt.Errorf("error creating discord session: %v", err)
	}

	b := NewBot(C, dg)
	b.Init(handler.NewHandlers(C), command.NewHandlers(C))

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	if err = dg.Open(); err != nil {
		return fmt.Errorf("error opening connection: %v", err)
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
	C = config.NewConfig()

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
