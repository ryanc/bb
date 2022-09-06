package commands

import (
	"fmt"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib/weather"

	log "github.com/sirupsen/logrus"
)

type WeatherHandler struct {
	Config bot.Config
	Name   string
}

func NewWeatherHandler(s string) *WeatherHandler {
	return &WeatherHandler{Name: s}
}

func (h *WeatherHandler) SetConfig(config bot.Config) {
	h.Config = config
}

func WeatherCommand(cmd *bot.Command, args []string) error {
	var (
		err error
		loc string
		w   weather.Weather
	)

	if len(args) != 1 {
		cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, "help: `!weather <CITY>,<STATE>,<COUNTRY>`")
		return nil
	}

	loc = args[0]

	if cmd.Config.OpenWeatherMapToken == "" {
		log.Error("OpenWeather token is not set")
		return nil
	}

	wc := weather.NewClient(cmd.Config.OpenWeatherMapToken)

	log.Debugf("weather requested for '%s'", loc)

	w, err = wc.Get(loc)
	if err != nil {
		log.Errorf("weather client error: %v", err)
		return nil
	}

	log.Debugf("weather returned for '%s': %+v", loc, w)

	cmd.Session.ChannelMessageSend(cmd.Message.ChannelID, fmt.Sprintf(
		"%s (%.1f, %.1f) — C:%.1f F:%.1f K:%.1f",
		loc,
		w.Coord.Lat,
		w.Coord.Lon,
		w.Main.Temp.Celcius(),
		w.Main.Temp.Fahrenheit(),
		w.Main.Temp.Kelvin(),
	))

	return nil
}
