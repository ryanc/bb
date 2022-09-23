package config

const (
	defaultPrefix = "!"
)

var (
	defaultReactions []string = []string{
		"👍", "🌶️", "🤣", "😂", "🍆", "🍑", "❤️", "💦", "😍", "💩",
		"🔥", "🍒", "🎉", "🥳", "🎊", "📉", "📈", "💀", "☠️",
	}
)

type (
	Config struct {
		Debug               bool           `mapstructure:"debug"`
		Handler             HandlerConfig  `mapstructure:"handler"`
		Prefix              string         `mapstructure:"prefix"`
		DiscordToken        string         `mapstructure:"discord_token"`
		OpenWeatherMapToken string         `mapstructure:"open_weather_map_token"`
		Mongo               MongoConfig    `mapstructure:"mongo"`
		Redis               RedisConfig    `mapstructure:"redis"`
		Postgres            PostgresConfig `mapstructure:"postgres"`
	}

	HandlerConfig struct {
		Reaction ReactionConfig `mapstructure:"reaction"`
	}

	ReactionConfig struct {
		Emojis   []string
		Channels []string
	}

	MongoConfig struct {
		Uri      string `mapstructure:"uri"`
		Database string `mapstructure:"database"`
	}

	RedisConfig struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"database"`
	}

	PostgresConfig struct {
		Uri string `mapstructure:"uri"`
	}
)

func NewConfig() *Config {
	var c *Config = &Config{}

	c.Prefix = defaultPrefix
	c.Handler.Reaction.Emojis = defaultReactions

	return c
}
