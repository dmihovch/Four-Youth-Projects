package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {
	signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	slackAPI := slack.New(botToken)

}
