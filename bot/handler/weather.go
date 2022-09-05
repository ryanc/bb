package handler

import (
	"fmt"
	"strings"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib/weather"
	"github.com/bwmarrin/discordgo"

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

func (h *WeatherHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		err error
		loc string
		w   weather.Weather
	)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!weather") {
		return
	}

	x := strings.SplitN(m.Content, " ", 2)

	if len(x) != 2 {
		s.ChannelMessageSend(m.ChannelID, "help: `!weather <CITY>,<STATE>,<COUNTRY>`")
		return
	}

	loc = x[1]

	if h.Config.OpenWeatherMapToken == "" {
		log.Error("OpenWeather token is not set")
		return
	}

	wc := weather.NewClient(h.Config.OpenWeatherMapToken)

	log.Debugf("weather requested for '%s'", loc)

	w, err = wc.Get(loc)
	if err != nil {
		log.Errorf("weather client error: %v", err)
		return
	}

	log.Debugf("weather returned for '%s': %+v", loc, w)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
		"%s (%.1f, %.1f) — C:%.1f F:%.1f K:%.1f",
		loc,
		w.Coord.Lat,
		w.Coord.Lon,
		w.Main.Temp.Celcius(),
		w.Main.Temp.Fahrenheit(),
		w.Main.Temp.Kelvin(),
	))
}
