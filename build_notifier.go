package functions

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/storage"
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
	LogsBucket string `json:"logsBucket,omitempty"`
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
	slack := notify.NewNotify(token, channel)
	err = slack.DirectNotify(string(build.Status))
	if err != nil {
		log.Fatal(err)
	}

	logs, err := getLogsFromGCS(ctx, build.LogsBucket, build.BuildID)
	if err != nil {
		log.Fatalf("fail to get logs from GCS %v: %v", build.LogsBucket, err)
	}
	if logs == nil {
		return nil
	}

	if err := slack.DirectNotify(string(logs)); err != nil {
		log.Fatalf("fail to slack posting: %v", err)
	}

	return nil
}

func getLogsFromGCS(ctx context.Context, bucket, id string) ([]byte, error) {
	log := "log-" + id + ".txt"

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(bucket, "gs://") {
		bucket = bucket[len("gs://"):]
	}
	rc, err := client.Bucket(bucket).Object(log).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	res, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return res, nil
}
