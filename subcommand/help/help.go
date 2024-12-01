package help

import (
	"github.com/ryo-kagawa/go-utils/commandline"
)

type HelpCommand struct{}

var _ = (commandline.SubCommand)(HelpCommand{})

func (HelpCommand) Name() string {
	return "help"
}
