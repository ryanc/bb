package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"git.kill0.net/chill9/beepboop/bot"
	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

const (
	OpenWeatherMapURI = "https://api.openweathermap.org"
)

var (
	EndpointWeather = lib.BuildURI(OpenWeatherMapURI, "/data/2.5/weather")
)

type (
	WeatherHandler struct {
		Config bot.Config
	}

	Temperature float32

	Weather struct {
		Main struct {
			Temp      Temperature `json:"temp"`
			FeelsLike Temperature `json:"feels_like"`
			TempMin   Temperature `json:"temp_min"`
			TempMax   Temperature `json:"temp_max"`
			Pressure  float32     `json:"pressure"`
			Humidity  float32     `json:"humidity"`
		} `json:"main"`

		Coord struct {
			Lon float32 `json:"lon"`
			Lat float32 `json:"lat"`
		} `json:"coord"`

		Rain struct {
			H1 float32 `json:"1h"`
			H3 float32 `json:"3h"`
		} `json:"rain"`
	}

	WeatherError struct {
		Message string `json:"message"`
	}
)

func (t *Temperature) Kelvin() float32 {
	return float32(*t)
}

func (t *Temperature) Fahrenheit() float32 {
	return ((float32(*t) - 273.15) * (9.0 / 5)) + 32
}

func (t *Temperature) Celcius() float32 {
	return float32(*t) - 273.15
}

func NewWeatherHandler() *WeatherHandler {
	return new(WeatherHandler)
}

func (h *WeatherHandler) SetConfig(config bot.Config) {
	h.Config = config
}

func (h *WeatherHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		loc  string
		w    Weather
		werr WeatherError
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

	req, err := http.NewRequest("GET", EndpointWeather, nil)
	if err != nil {
		log.Errorf("failed to create new request: %s", err)
		return
	}

	q := req.URL.Query()
	q.Add("q", loc)
	q.Add("appid", h.Config.OpenWeatherMapToken)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("HTTP request failed: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("reading HTTP response failed: %s", err)
		return
	}

	if resp.StatusCode != 200 {
		err = json.Unmarshal(body, &werr)
		if err != nil {
			log.Debugf("%s\n", body)
			log.Errorf("unmarshaling JSON failed: %s", err)
			return
		}
		log.Warnf("error: (%s) %s", resp.Status, werr.Message)
		return
	}

	log.Debugf("weather requested for '%s'\n", loc)

	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Debugf("%s\n", body)
		log.Errorf("unmarshaling JSON failed: %s", err)
		return
	}

	log.Debugf("weather returned for '%s': %+v\n", loc, w)

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
