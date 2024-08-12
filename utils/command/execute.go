package command

func Execute(command Command, subcommandList []Subcommand, args []string) (string, error) {
	if len(args) >= 1 {
		for _, subcommand := range subcommandList {
			if subcommand.Name() == args[0] {
				return subcommand.Execute()
			}
		}
	}
	return command.Execute()
}
