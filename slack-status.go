package main

import (
	"github.com/docopt/docopt-go"
	"github.com/fhuitelec/slack-status/slack"
)

var arguments map[string]interface{}

func init() {

	usage := `Slack status editor.

Usage:
  slack-status --emoji=<emoji> --status=<status>
  slack-status --status=<status>
  slack-status --reset-status
  slack-status --having-lunch
  slack-status -h | --help

Options:
  -h --help         Show this screen.
  --emoji=<emoji>   Emoji to use in your Slack status.
  --status=<status> Text to use in your Slack status.
  --reset-status   Reset your Slack status`

	arguments, _ = docopt.Parse(usage, nil, true, "Slack status", false)
}

func main() {
	status, emoji := processArguments(arguments)

	slack.ChangeProfileStatus(status, emoji)
}

func processArguments(arguments map[string]interface{}) (string, string) {
	status, ok := arguments["--status"].(string)
	if !ok {
		status = ""
	}

	emoji, ok := arguments["--emoji"].(string)
	if !ok {
		emoji = ""
	}

	resetStatus, ok := arguments["--reset-status"].(bool)
	if !ok {
		resetStatus = false
	}

	havingLunch, ok := arguments["--having-lunch"].(bool)
	if !ok {
		resetStatus = false
	}

	if havingLunch {
		status, emoji = "Having lunch", ":poultry_leg:"
	}

	if resetStatus {
		status, emoji = "", ""
	}

	return status, emoji
}
