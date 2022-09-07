package main

import (
	"flag"
	"fmt"

	//"log"

	"os"
	"os/signal"
	"syscall"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/bot/handler"
	"git.kill0.net/chill9/beepboop/lib"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	C bot.Config

	handlers []bot.MessageCreateHandler = []bot.MessageCreateHandler{
		handler.NewReactionHandler(),
	}
)

func main() {
	setupConfig()

	lib.SeedMathRand()

	if C.DiscordToken == "" {
		log.Fatalf("Discord token is not set")
	}

	dg, err := discordgo.New(fmt.Sprintf("Bot %s", C.DiscordToken))
	if err != nil {
		log.Fatalf("error creating Discord session: %v\n", err)
	}

	for _, h := range handlers {
		h.SetConfig(C)
		dg.AddHandler(h.Handle)
	}

	b := bot.NewBot(dg, C)
	b.RegisterCommands()

	dg.AddHandler(bot.NewCommandHandler(b))

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
}

func setupConfig() {
	var err error

	C = bot.NewConfig()

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
