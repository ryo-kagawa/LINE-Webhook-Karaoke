package help

import (
	"strings"

	"github.com/ryo-kagawa/go-utils/commandline"
)

type HelpCommand struct{}

var _ = (commandline.SubCommand)(HelpCommand{})

func (HelpCommand) Execute(arguments []string) (string, error) {
	return strings.Trim(
		`
usage:
  LINE-Wehbhook-Karaoke [subcommand]

subcommands:
  initialize Initialize Database
	version    Various versions

usage: LINE-Wehbhook-Karaoke
environment:
  DATABASE_PASSWORD
  DATABASE_URL
  DATABASE_USER
  LINE_CHANNEL_SECRET
  LINE_CHANNEL_TOKEN

usage: LINE-Wehbhook-Karaoke intialize
environment:
  DATABASE_PASSWORD
  DATABASE_URL
  DATABASE_USER
usage: LINE-Wehbhook-Karaoke version
`,
		"\n",
	), nil
}

func (HelpCommand) Name() string {
	return "help"
}
