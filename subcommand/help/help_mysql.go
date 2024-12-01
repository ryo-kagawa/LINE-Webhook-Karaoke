//go:build mysql

package help

import "strings"

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
  DATABASE_MYSQL_ADDRESS
  DATABASE_MYSQL_USER
  DATABASE_MYSQL_PASSWORD
  DATABASE_MYSQL_DATABASE
  LINE_CHANNEL_SECRET
  LINE_CHANNEL_TOKEN

usage: LINE-Wehbhook-Karaoke intialize
environment:
  DATABASE_MYSQL_ADDRESS
  DATABASE_MYSQL_USER
  DATABASE_MYSQL_PASSWORD
  DATABASE_MYSQL_DATABASE
usage: LINE-Wehbhook-Karaoke version
`,
		"\n",
	), nil
}
