package command

import (
	"fmt"

	"git.kill0.net/chill9/beepboop/lib/weather"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) Weather(args []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	var (
		err error
		loc string
		w   weather.Weather
	)

	if len(args) != 1 {
		s.ChannelMessageSend(m.ChannelID, "help: `!weather <CITY>,<STATE>,<COUNTRY>`")
		return nil
	}

	loc = args[0]

	if h.config.OpenWeatherMapToken == "" {
		log.Error("OpenWeather token is not set")
		return nil
	}

	wc := weather.NewClient(h.config.OpenWeatherMapToken)

	log.Debugf("weather requested for '%s'", loc)

	w, err = wc.Get(loc)
	if err != nil {
		log.Errorf("weather client error: %v", err)
		return nil
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

	return nil
}
