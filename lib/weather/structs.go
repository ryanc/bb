package weather

type (
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

		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
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
