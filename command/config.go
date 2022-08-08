package command

const (
	defaultPrefix = "!"
)

var (
	defaultReactions []string = []string{"👍", "🌶️", "🤣", "😂", "🍆", "🍑", "❤️", "💦", "😍", "💩", "🔥", "🍒", "🎉", "🥳", "🎊"}
)

type (
	Config struct {
		Handler HandlerConfig `mapstructure:"handler"`
		Prefix  string        `mapstructure:"prefix"`
	}

	HandlerConfig struct {
		Reaction ReactionConfig `mapstructure:"reaction"`
		Weather  WeatherConfig  `mapstructure:"weather"`
	}

	ReactionConfig struct {
		Emojis   []string
		Channels []string
	}
)

func NewConfig() Config {
	var c Config

	c.Prefix = defaultPrefix
	c.Handler.Reaction.Emojis = defaultReactions

	return c
}
