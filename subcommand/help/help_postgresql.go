//go:build postgresql

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
  DATABASE_POSTGRESQL_HOST
  DATABASE_POSTGRESQL_PORT
  DATABASE_POSTGRESQL_USER
  DATABASE_POSTGRESQL_PASSWORD
  DATABASE_POSTGRESQL_DATABASE
  LINE_CHANNEL_SECRET
  LINE_CHANNEL_TOKEN

usage: LINE-Wehbhook-Karaoke intialize
environment:
  DATABASE_POSTGRESQL_HOST
  DATABASE_POSTGRESQL_PORT
  DATABASE_POSTGRESQL_USER
  DATABASE_POSTGRESQL_PASSWORD
  DATABASE_POSTGRESQL_DATABASE
usage: LINE-Wehbhook-Karaoke version
`,
		"\n",
	), nil
}
