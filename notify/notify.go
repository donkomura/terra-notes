package notify

import (
	"github.com/donkomura/terra-notes/parser"
	"github.com/nlopes/slack"
)

type notifyConfig struct {
	slackAPIToken string
	slackChannel  string
}

func NewNotify(slackToken, slackChannel string) *notifyConfig {
	return &notifyConfig{
		slackAPIToken: slackToken,
		slackChannel:  slackChannel,
	}
}

// return message's timestamp to craete therads
func (c *notifyConfig) postMessage(text string) (string, error) {
	_, ts, err := slack.New(c.slackAPIToken).PostMessage(
		c.slackChannel,
		slack.MsgOptionText("```\n"+text+"\n```", false),
	)

	return ts, err
}

func (c *notifyConfig) postThread(ts string, text string) error {
	_, _, err := slack.New(c.slackAPIToken).PostMessage(
		c.slackChannel,
		slack.MsgOptionText("```\n"+text+"\n```", false),
		slack.MsgOptionTS(ts),
	)

	return err
}

func (c *notifyConfig) Notify(parsedBody parser.PlanResult) error {
	summary := parsedBody.Summary
	if parsedBody.Status == 1 {
		summary = parsedBody.Error
	}
	ts, err := c.postMessage(summary)
	if err != nil {
		return err
	}

	for _, r := range parsedBody.Plan {
		if err := c.postThread(ts, r); err != nil {
			return err
		}
	}

	return nil
}

func (c *notifyConfig) DirectNotify(msg string) error {
	_, err := c.postMessage(msg)
	return err
}
