package main

import (
	"fmt"
	"os"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/rootcommand"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/subcommand/help"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/subcommand/initialize"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/subcommand/version"
	"github.com/ryo-kagawa/go-utils/commandline"
)

func main() {
	result, err := commandline.Execute(
		rootcommand.Command{},
		os.Args[1:],
		help.HelpCommand{},
		version.Version{},
		initialize.Initialize{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
