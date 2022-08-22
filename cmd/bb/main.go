package main

import (
	"flag"
	"fmt"

	//"log"

	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"git.kill0.net/chill9/beepboop/command"
	"git.kill0.net/chill9/beepboop/lib"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	Token string

	defaultReactions []string = []string{"👍", "🌶️", "🤣", "😂", "🍆", "🍑", "❤️", "💦", "😍", "💩", "🔥", "🍒", "🎉", "🥳", "🎊"}

	C command.Config

	handlers []command.CommandHandler = []command.CommandHandler{
		command.NewCoinHandler(),
		command.NewPingHandler(),
		command.NewRollHandler(),
		command.NewRouletteHandler(),
		command.NewTimeHandler(),
		command.NewVersionHandler("version"),
		command.NewWeatherHandler(),
	}
)

func main() {
	var (
		err error
	)

	flag.Bool("debug", false, "enable debug logging")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	lib.SeedMathRand()

	viper.SetDefault("handler.reaction.emojis", defaultReactions)
	viper.SetEnvPrefix("BEEPBOOP")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()

	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		log.Fatalf("fatal error config file: %v", err)
	}

	Token, ok := viper.Get("discord_token").(string)

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	if Token == "" {
		log.Fatalf("Discord token is not set")
	}

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	dg, err := discordgo.New(fmt.Sprintf("Bot %s", Token))
	if err != nil {
		log.Fatalf("error creating Discord session: %v\n", err)
	}

	for _, h := range handlers {
		h.SetConfig(C)
		dg.AddHandler(h.Handle)
	}

	dg.AddHandler(reactionHandler)
	dg.AddHandler(praiseHandler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

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
}

func praiseHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "good bot") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> Thank you, daddy.", m.Author.ID))
	}
}

func reactionHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	emojis := C.Handler.Reaction.Emojis
	channels := C.Handler.Reaction.Channels

	if len(emojis) == 0 {
		log.Warning("emoji list is empty")
		return
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Fatalf("unable to get channel name: %v", err)
	}

	if len(channels) > 0 && !contains(channels, channel.Name) {
		return
	}

	for _, a := range m.Attachments {
		if strings.HasPrefix(a.ContentType, "image/") {
			r := emojis[rand.Intn(len(emojis))]
			s.MessageReactionAdd(m.ChannelID, m.ID, r)
		}
	}

	for range m.Embeds {
		r := emojis[rand.Intn(len(emojis))]
		s.MessageReactionAdd(m.ChannelID, m.ID, r)
	}

}

func contains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}
