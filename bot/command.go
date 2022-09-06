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
		Name    string
		Config  Config
		Func    func(cmd *Command, args []string) error
		Session *discordgo.Session
		Message *discordgo.MessageCreate
	}
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

func NewCommandHandler(config Config) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		var cmd *Command

		if m.Author.ID == s.State.User.ID {
			return
		}

		if !lib.HasCommand(m.Content, config.Prefix) {
			return
		}

		cmdName, args := lib.SplitCommandAndArgs(m.Content, config.Prefix)

		cmd, ok := GetCommand(cmdName)
		if ok {
			cmd.Config = config
			cmd.Name = cmdName
			cmd.Session = s
			cmd.Message = m

			log.Debugf("command: %+v, args: %+v", cmd.Name, args)
			cmd.Func(cmd, args)

			return
		}

		log.Warnf("unknown command: %+v, args: %+v", cmdName, args)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("unknown command: %s", cmdName))
	}
}
