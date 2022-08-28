package weather

import (
	"encoding/json"
	"errors"
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
	EndpointWeather        = lib.BuildURI(OpenWeatherMapURI, "/data/2.5/weather")
	ErrUnmarshal           = errors.New("unmarshaling JSON failed")
	ErrReadingResponse     = errors.New("reading HTTP response failed")
	ErrRequestFailed       = errors.New("HTTP request failed")
	ErrCreateRequestFailed = errors.New("failed to create new HTTP request")
)

func NewClient(token string) *WeatherClient {
	return &WeatherClient{token}
}

func (c *WeatherClient) Get(loc string) (w Weather, err error) {
	var werr WeatherError

	req, err := http.NewRequest("GET", EndpointWeather, nil)
	if err != nil {
		err = fmt.Errorf("%s: %s", ErrCreateRequestFailed, err)
		return
	}

	q := req.URL.Query()
	q.Add("q", loc)
	q.Add("appid", c.token)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("%s: %s", ErrRequestFailed, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("%s: %s", ErrReadingResponse, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = json.Unmarshal(body, &werr)
		if err != nil {
			log.Debugf("%s", body)
			err = fmt.Errorf("%s: %s", ErrUnmarshal, err)
			return
		}

		err = fmt.Errorf("error: (%s) %s", resp.Status, werr.Message)
		return
	}

	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Debugf("%s", body)
		err = fmt.Errorf("%s: %s", ErrUnmarshal, err)
		return
	}

	return
}
