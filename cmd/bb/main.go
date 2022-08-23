package main

import (
	"flag"
	"fmt"

	//"log"

	"os"
	"os/signal"
	"syscall"

	"git.kill0.net/chill9/beepboop/bot"
	handler "git.kill0.net/chill9/beepboop/bot/handlers"
	"git.kill0.net/chill9/beepboop/command"
	"git.kill0.net/chill9/beepboop/lib"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	C bot.Config

	handlers []command.CommandHandler = []command.CommandHandler{
		command.NewCoinHandler(),
		command.NewPingHandler(),
		command.NewRollHandler(),
		command.NewRouletteHandler(),
		command.NewTimeHandler(),
		command.NewVersionHandler("version"),
		command.NewWeatherHandler(),
		handler.NewReactionHandler(),
	}
)

func main() {
	var (
		err error
	)

	C = bot.NewConfig()

	flag.Bool("debug", false, "enable debug logging")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	lib.SeedMathRand()

	viper.SetEnvPrefix("BEEPBOOP")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	viper.BindEnv("DISCORD_TOKEN")
	viper.BindEnv("OPEN_WEATHER_MAP_TOKEN")

	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		log.Fatalf("fatal error config file: %v", err)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

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
