package main

import (
	"fmt"
	"os"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/subcommand/initialize"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/subcommand/version"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/utils/command"
)

func main() {
	result, err := command.Execute(
		Command{},
		[]command.Subcommand{
			initialize.Initialize{},
			version.Version{},
		},
		os.Args[1:],
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
