package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	//Slack Client
	client := slack.New(token, slack.OptionDebug(true))

	//Attachement to send to channel
	attachment := slack.Attachment{
		Pretext: "Message From Bot",
		Text:    "Hello, I am a Slack Bot",
		Color:   "#fff",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}

	_, timestamp, err := client.PostMessage(
		channelID,
		slack.MsgOptionText("Check this message from bot", false),
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Message was sent at" + timestamp)
}
