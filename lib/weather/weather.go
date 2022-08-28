package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"git.kill0.net/chill9/beepboop/lib"

	log "github.com/sirupsen/logrus"
)

type (
	WeatherClient struct {
		token string
	}
)

const (
	OpenWeatherMapURI = "https://api.openweathermap.org"
)

var (
	EndpointWeather = lib.BuildURI(OpenWeatherMapURI, "/data/2.5/weather")
)

func NewClient(token string) *WeatherClient {
	return &WeatherClient{token}
}

func (c *WeatherClient) Get(loc string) (w Weather, err error) {
	var (
		werr WeatherError
	)

	req, err := http.NewRequest("GET", EndpointWeather, nil)
	if err != nil {
		err = fmt.Errorf("failed to create new request: %s", err)
		return
	}

	q := req.URL.Query()
	q.Add("q", loc)
	q.Add("appid", c.token)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("HTTP request failed: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("reading HTTP response failed: %s", err)
		return
	}

	if resp.StatusCode != 200 {
		err = json.Unmarshal(body, &werr)
		if err != nil {
			log.Debugf("%s", body)
			err = fmt.Errorf("unmarshaling JSON failed: %s", err)
			return
		}

		err = fmt.Errorf("error: (%s) %s", resp.Status, werr.Message)
		return
	}

	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Debugf("%s", body)
		log.Errorf("unmarshaling JSON failed: %s", err)
		return
	}

	return
}
