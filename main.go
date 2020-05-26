package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

const (
	envSlackWebhook    = "SLACK_WEBHOOK"
	envSlackChannel    = "SLACK_CHANNEL"
	envSlackMessage    = "SLACK_MESSAGE"
	envSlackIcon       = "SLACK_ICON"
	envSlackIconEmoji  = "SLACK_ICON_EMOJI"
	envSlackTitle      = "SLACK_TITLE"
	envSlackColor      = "SLACK_COLOR"
	envGithubRef       = "GITHUB_REF"
	envGithubEventName = "GITHUB_EVENT_NAME"
	envGithubRepo      = "GITHUB_REPOSITORY"
	envGithubSha       = "GITHUB_SHA"
	envGithubActor     = "GITHUB_ACTOR"
)

func main() {
	endpoint := getEnvOrDefault(envSlackWebhook, "")
	channel := getEnvOrDefault(envSlackChannel, "")
	icon := getEnvOrDefault(envSlackIcon, "")
	iconEmoji := getEnvOrDefault(envSlackIconEmoji, "Ghost")
	title := getEnvOrDefault(envSlackTitle, "")
	message := getEnvOrDefault(envSlackMessage, "")
	actor := getEnvOrDefault(envGithubActor, "")
	ref := getEnvOrDefault(envGithubRef, "")
	eventName := getEnvOrDefault(envGithubEventName, "")
	if len(endpoint) == 0 {
		fmt.Printf("%v must be defined \n", envSlackWebhook)
		os.Exit(1)
	}

	if len(channel) == 0 {
		fmt.Printf("%v must be defined \n", envSlackChannel)
		os.Exit(1)
	}

	fields := []slack.AttachmentField{
		{
			Value: fmt.Sprintf("*Repo*: "+envGithubRepo+" *Ref*: %v %v  <https://github.com/"+envGithubRepo+"/commit/"+os.Getenv("GITHUB_SHA")+"/checks\"|Details>", ref, eventName),
			Short: false,
		},

		{
			Title: title,
			Value: message,
			Short: false,
		},
	}
	attachment := slack.Attachment{
		AuthorName: actor,
		Color:      getEnvOrDefault(envSlackColor, "good"),
		AuthorLink: "http://github.com/" + actor,
		AuthorIcon: "http://github.com/" + actor + ".png?size=32",
		Fields:     fields,
	}

	webhookMessage := slack.WebhookMessage{
		Channel:   channel,
		IconEmoji: iconEmoji,
		IconURL:   icon,
		Username:  "Github Alerter",
		Attachments: []slack.Attachment{
			attachment,
		},
	}
	fmt.Println("Sending to Slack Channel")
	slack.PostWebhook(endpoint, &webhookMessage)
}

func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}
