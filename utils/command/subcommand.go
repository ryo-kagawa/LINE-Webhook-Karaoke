package command

type Subcommand interface {
	Command
	Name() string
}
