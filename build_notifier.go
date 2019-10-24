package functions

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/donkomura/terra-notes/notify"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type BuildResult struct {
	BuildID    string `json:"id,omitempty"`
	ProjectID  string `json:"projectId,omitempty"`
	Status     string `json:"status,omitempty"`
	StartTime  string `json:"startTime,omitempty"`
	FinishTime string `json:"finishTime,omitempty"`
	LogURL     string `json:"logUrl,omitempty"`
}

func BuildStatus(ctx context.Context, m PubSubMessage) error {
	var build BuildResult

	err := json.Unmarshal(m.Data, &build)
	if err != nil {
		log.Fatalf("fail to load json: %v", err)
	}

	token := os.Getenv("SLACK_API_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL_NAME")

	if token == "" || channel == "" {
		log.Fatalf("invalid slack token or channel: token[%v], channel[%v]", token, channel)
	}

	if err := notify.NewNotify(token, channel).DirectNotify(string(m.Data)); err != nil {
		log.Fatalf("fail to slack posting: %v", err)
	}

	return nil
}
