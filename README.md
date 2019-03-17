## Slack status CLI

This bit of CLI allows you to change your status using your command line.

## Installation

### Homebrew

_See [`Linuxbrew`](https://docs.brew.sh/Homebrew-on-Linux) for linux platforms._

```shell
# This will add my personal tap fhuitelec/fhuitelec
brew install fhuitelec/fhuitelec/slack-status
```

### Binary

You can download the binary from the [release page](https://github.com/fhuitelec/slack-status/releases) for your platform.

## Usage

Basic usage:

```
# Add a new Slack status using an emoji and status
slack-status --emoji=":coffee:" --status="Coffee break"
# You can specify your status only, with an emoji, in which case Slack will add its own emoji (currently :speech_ballon:)
slack-status --status="Doing something"
# Resets your Slack status with an empty status & no emoji
```

You will be prompted and be asked your Slack token, you can find it or issue it [here](https://api.slack.com/custom-integrations/legacy-tokens#legacy_token_generator) (make sure you have an on-going connected Slack session - i.e. you are connected to Slack).

### Configuration

Instead of adding your Slack token interactively, you can create the configuration file yourself in `~/config/slack-status/config.json` and add the token:

```json
{
    "token": "your-token-here"
}
```

### Emojis

Since Slack handles emojis though unique text IDs, you can use standard Slack emojis as well as your Organization's own emojis.

## Contributing

### Requirements

- golang: last tested on version `1.12.1`

### Build

```shell
# Get it
go get github.com/fhuitelec/slack-status
cd $GOPATH/src/github.com/fhuitelec/slack-status
# Built it
GOOS=linux GOARCH=amd64 go build -o $PWD/build/slack-status-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o $PWD/build/slack-status-darwin-amd64
```

## Known limitations

### No tests

I'm new to the golang language & ecosystem and was more driven by delivery than quality, sorry 🤷‍♂️

### No clear vendor dependencies policy

Same as above: I'm new to the golang language & ecosystem 🙄

### One slack profile only

This CLI can only handle one Slack profile at a time, not multiple.

### Legacy Slack token

The current implementation is based on the Slack's legacy token system, there's no ongoing work to change that behaviour, feel free to create a pull request.
