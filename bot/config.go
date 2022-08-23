package bot

const (
	defaultPrefix = "!"
)

var (
	defaultReactions []string = []string{"👍", "🌶️", "🤣", "😂", "🍆", "🍑", "❤️", "💦", "😍", "💩", "🔥", "🍒", "🎉", "🥳", "🎊"}
)

type (
	Config struct {
		Handler             HandlerConfig `mapstructure:"handler"`
		Prefix              string        `mapstructure:"prefix"`
		DiscordToken        string        `mapstructure:"discord_token"`
		OpenWeatherMapToken string        `mapstructure:"open_weather_map_token"`
	}

	HandlerConfig struct {
		Reaction ReactionConfig `mapstructure:"reaction"`
		Weather  WeatherConfig  `mapstructure:"weather"`
	}

	ReactionConfig struct {
		Emojis   []string
		Channels []string
	}

	WeatherConfig struct {
		Token string `mapstructure:"token"`
	}
)

func NewConfig() Config {
	var c Config

	c.Prefix = defaultPrefix
	c.Handler.Reaction.Emojis = defaultReactions

	return c
}
