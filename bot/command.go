package bot

import (
	"fmt"

	"git.kill0.net/chill9/beepboop/lib"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	DefaultCommander *Commander
)

type (
	Commander struct {
		commands map[string]*Command
	}

	Command struct {
		Name   string
		Config Config
		Func   CommandFunc
		NArgs  int
	}

	CommandFunc func(args []string, m *discordgo.MessageCreate) error
)

func init() {
	DefaultCommander = NewCommander()
}

func NewCommander() *Commander {
	cmdr := new(Commander)
	cmdr.commands = make(map[string]*Command)
	return cmdr
}

func (cmdr *Commander) AddCommand(cmd *Command) {
	cmdr.commands[cmd.Name] = cmd
}

func (cmdr *Commander) GetCommand(name string) (*Command, bool) {
	cmd, ok := cmdr.commands[name]
	return cmd, ok
}

func AddCommand(cmd *Command) {
	DefaultCommander.AddCommand(cmd)
}

func GetCommand(name string) (*Command, bool) {
	cmd, ok := DefaultCommander.GetCommand(name)
	return cmd, ok
}

func NewCommandHandler(bot *Bot) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		var cmd *Command

		if m.Author.ID == s.State.User.ID {
			return
		}

		if !lib.HasCommand(m.Content, bot.Config.Prefix) {
			return
		}

		cmdName, arg := lib.SplitCommandAndArg(m.Content, bot.Config.Prefix)

		cmd, ok := GetCommand(cmdName)

		args := lib.SplitArgs(arg, cmd.NArgs)

		if ok {
			cmd.Config = bot.Config

			log.Debugf("command: %v, args: %v, nargs: %d", cmd.Name, args, len(args))
			cmd.Func(args, m)

			return
		}

		log.Warnf("unknown command: %v, args: %v, nargs: %d", cmdName, args, len(args))
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("unknown command: %s", cmdName))
	}
}
